package parse

import (
	"fmt"
	"go/ast"
	"go/token"
	"path"
	"sync"

	"github.com/becheran/smock/internal/logger"
	"github.com/becheran/smock/internal/model"
	"github.com/becheran/smock/internal/pathhelper"
	"golang.org/x/exp/slices"
)

func ParseInterfaceInPackage(pkg *ast.Package, interfaceName string) (i model.InterfaceResult, err error) {
	logger.Printf("Parse interface '%s' in package '%s'", interfaceName, pkg.Name)

	for path, file := range pkg.Files {
		i, err = ParseInterfaceInFile(file, interfaceName, path)
		if err == nil {
			return i, nil
		}
	}
	return model.InterfaceResult{}, fmt.Errorf("interface '%s' not found in package '%s'", interfaceName, pkg.Name)
}

func ParseInterfaceInFile(file *ast.File, interfaceName, path string) (i model.InterfaceResult, err error) {
	logger.Printf("Parse interface %s in file '%s'", interfaceName, file.Name)
	for _, decl := range file.Decls {
		ts, err := getTypeSpec(decl)
		if err != nil {
			continue
		}
		if ts.Name.Name == interfaceName {
			return parseInterface(ts, file.Name.Name, path, file.Imports)
		}
	}
	return model.InterfaceResult{}, fmt.Errorf("interface %s not found in file %s", interfaceName, file.Name)
}

func ParseInterfaceAtPosition(fset *token.FileSet, file *ast.File, startLine int) (i model.InterfaceResult, err error) {
	logger.Printf("Parse interface in file '%s:%d'", file.Name, startLine)

	for _, decl := range file.Decls {
		line := fset.Position(decl.Pos()).Line

		if line <= startLine {
			continue
		}

		ts, err := getTypeSpec(decl)
		if err != nil {
			return model.InterfaceResult{}, err
		}

		return parseInterface(ts, file.Name.Name, "./", file.Imports)
	}

	return model.InterfaceResult{}, fmt.Errorf("interface at %s:%d not found", file.Name, startLine)
}

func parseInterface(ts *ast.TypeSpec, pkgName, file string, imports []*ast.ImportSpec) (i model.InterfaceResult, err error) {
	logger.Printf("Parse interface '%s' in file '%s'", ts.Name.Name, file)
	dir := path.Dir(pathhelper.PathToUnix(file))

	i.Name = ts.Name.Name
	i.PackageName = pkgName

	interfaceType, ok := ts.Type.(*ast.InterfaceType)
	if !ok {
		if ref := expToReference(ts.Type, pkgName); ref != nil {
			packageId := ref.PackageID
			pkg, err := parsePackage(packageId, imports, dir)
			if err != nil {
				return model.InterfaceResult{}, fmt.Errorf("failed to resolve package reference. %w", err)
			}
			res, err := ParseInterfaceInPackage(pkg, ref.Name)
			if err != nil {
				return model.InterfaceResult{}, fmt.Errorf("failed to parse referenced interface '%s'. %w", i.Name, err)
			}
			res.Name = i.Name
			res.PackageName = i.PackageName
			return res, nil
		}
		return model.InterfaceResult{}, fmt.Errorf("unexpected type %T", ts.Type)
	}

	if interfaceType.Methods == nil {
		return model.InterfaceResult{}, fmt.Errorf("unexpected empty interface")
	}

	identResolver := identResolver{PackageName: i.PackageName, UsedImports: make(map[string]struct{})}
	if ts.TypeParams != nil {
		i.Types = identResolver.fieldListToIdent(ts.TypeParams.List)
	}
	generics := make(map[string]struct{})
	for _, genType := range i.Types {
		generics[genType.Name] = struct{}{}
	}
	identResolver.Generics = generics

	referencedInterfaces := []*model.Reference{}
	for _, it := range interfaceType.Methods.List {
		if ref := expToReference(it.Type, pkgName); ref != nil {
			logger.Printf("Found referenced interface '%s' in '%s'", ref.Name, ref.PackageID)
			referencedInterfaces = append(referencedInterfaces, ref)
			continue
		}
		if len(it.Names) != 1 {
			continue
		}
		name := it.Names[0]
		if !name.IsExported() {
			continue
		}
		logger.Printf("found exported method '%s'", name)
		switch meth := it.Type.(type) {
		case *ast.FuncType:
			getList := func(list *ast.FieldList) []*ast.Field {
				if list == nil {
					return nil
				}
				return list.List
			}
			method := model.Method{
				Name:       name.String(),
				TypeParams: identResolver.fieldListToIdent(getList(meth.TypeParams)),
				Params:     identResolver.fieldListToIdent(getList(meth.Params)),
				Results:    identResolver.fieldListToIdent(getList(meth.Results)),
			}
			i.Methods = append(i.Methods, method)
		default:
			return model.InterfaceResult{}, fmt.Errorf("unexpected type expression %T", it.Type)
		}
	}

	// TODO: Move to own go function
	var wg sync.WaitGroup
	var mux sync.Mutex
	packages := make(map[string]*ast.Package)
	var packagesErr error
	for _, ref := range referencedInterfaces {
		packageId := ref.PackageID
		mux.Lock()
		if _, ok := packages[packageId]; !ok {
			packages[packageId] = nil

			wg.Add(1)
			go func() {
				defer wg.Done()
				pkg, err := parsePackage(packageId, imports, dir)
				mux.Lock()
				defer mux.Unlock()
				if err != nil {
					packagesErr = err
				} else {
					packages[packageId] = pkg
				}
			}()
		}
		mux.Unlock()
	}

	wg.Wait()
	if packagesErr != nil {
		return model.InterfaceResult{}, fmt.Errorf("failed to resolve package reference. %w", packagesErr)
	}

	inheritInterfaces := []*model.InterfaceResult{}
	for _, ref := range referencedInterfaces {
		packageId := ref.PackageID
		name := ref.Name

		wg.Add(1)
		go func() {
			defer wg.Done()
			res, err := ParseInterfaceInPackage(packages[packageId], name)
			mux.Lock()
			defer mux.Unlock()
			if err != nil {
				packagesErr = err
			} else {
				inheritInterfaces = append(inheritInterfaces, &res)
			}
		}()
	}

	wg.Wait()
	if packagesErr != nil {
		return model.InterfaceResult{}, fmt.Errorf("failed to resolve referenced interfaces for '%s'. %w", i.Name, packagesErr)
	}

	for usedImport := range identResolver.UsedImports {
		logger.Printf("Add used import '%s' to result", usedImport)

		if usedImport == i.PackageName {
			continue
		}
		var foundImport *model.Import
		for _, astImp := range imports {
			imp := model.ImportFromAst(astImp)
			if imp.ImportName() == usedImport {
				foundImport = &imp
				break
			}
		}
		if foundImport == nil {
			return model.InterfaceResult{}, fmt.Errorf("import '%s' not found", usedImport)
		}
		i.Imports = append(i.Imports, *foundImport)
	}

	slices.SortFunc(inheritInterfaces, func(a, b *model.InterfaceResult) bool { return a.Name < b.Name })
	for _, inheritInterface := range inheritInterfaces {
		i.Methods = append(i.Methods, inheritInterface.Methods...)

		for _, genImport := range inheritInterface.Imports {
			if slices.ContainsFunc(i.Imports, func(a model.Import) bool { return a.ImportName() == genImport.ImportName() }) {
				logger.Printf("Import %s already added in original interface", genImport.ImportName())
				continue
			}
			i.Imports = append(i.Imports, genImport)
		}
	}

	slices.SortFunc(i.Imports, func(a, b model.Import) bool { return a.ImportName() < b.ImportName() })

	return i, nil
}

func expToReference(exp ast.Expr, pkgName string) *model.Reference {
	switch meth := exp.(type) {
	case *ast.SelectorExpr:
		packageID := pkgName
		if xIdent, ok := meth.X.(*ast.Ident); ok {
			packageID = xIdent.String()
		}
		return &model.Reference{
			PackageID: packageID,
			Name:      meth.Sel.String(),
		}
	case *ast.Ident:
		return &model.Reference{
			Name: meth.String(),
		}
	}
	return nil
}

type identResolver struct {
	PackageName string
	UsedImports map[string]struct{}
	Generics    map[string]struct{}
}

func (f *identResolver) fieldListToIdent(list []*ast.Field) (res []model.Ident) {
	if list == nil {
		return
	}
	for _, l := range list {
		tr := typeResolver{
			PackageName: f.PackageName,
			UsedImports: f.UsedImports,
			Generics:    f.Generics,
		}
		identType := tr.resolveType(l.Type)

		if len(l.Names) == 0 {
			res = append(res, model.Ident{Type: identType})
		} else {
			for _, name := range l.Names {
				res = append(res, model.Ident{Name: name.Name, Type: identType})
			}
		}
	}
	return
}

type typeResolver struct {
	PackageName string
	UsedImports map[string]struct{}
	Generics    map[string]struct{}
}

func (tr *typeResolver) isGenericIdentifier(id string) bool {
	if tr.Generics == nil {
		return false
	}
	_, ok := tr.Generics[id]
	return ok
}

func (tr *typeResolver) resolveType(exp ast.Expr) (identType string) {
	if exp == nil {
		return "" // For example in case of nil
	}
	switch t := exp.(type) {
	case *ast.Ident:
		identType = t.String()
		if !tr.isGenericIdentifier(identType) {
			if t.IsExported() {
				identType = tr.PackageName + "." + identType
			}
			tr.UsedImports[tr.PackageName] = struct{}{}
		}
	case *ast.SelectorExpr:
		if xIdent, ok := t.X.(*ast.Ident); ok {
			identType = xIdent.String()
		}
		tr.UsedImports[identType] = struct{}{}
		identType += "." + t.Sel.String()
	case *ast.MapType:
		identType += "map["
		identType += tr.resolveType(t.Key)
		identType += "]"
		identType += tr.resolveType(t.Value)
	case *ast.ArrayType:
		identType += "[" + tr.resolveType(t.Len) + "]"
		identType += tr.resolveType(t.Elt)
	case *ast.Ellipsis:
		identType += "..."
	case *ast.FuncType:
		identType += "func("
		for _, param := range t.Params.List {
			for i := 0; i < len(param.Names); i++ {
				identType += tr.resolveType(param.Type)
				identType += ", "
			}
		}
		if len(t.Params.List) > 0 {
			identType = identType[:len(identType)-2]
		}
		identType += ")"
		if t.Results != nil && len(t.Results.List) > 0 {
			identType += " "
			isMultiple := len(t.Results.List) > 1 || len(t.Results.List[0].Names) > 1
			if isMultiple {
				identType += "("
			}
			for _, param := range t.Results.List {
				for i := 0; i < len(param.Names); i++ {
					identType += tr.resolveType(param.Type)
					identType += ", "
				}
			}
			if isMultiple {
				identType = identType[:len(identType)-2] + ")"
			}
		}
	case *ast.StructType:
		identType += "struct{"
		for idx, field := range t.Fields.List {
			identType += " "
			for nameIdx, name := range field.Names {
				identType += name.String() + " "
				if nameIdx+1 < len(field.Names) {
					identType += ","
				}
			}
			identType += tr.resolveType(field.Type) + " "
			if idx+1 < len(t.Fields.List) {
				identType += ","
			}
		}
		identType += "}"
	case *ast.ChanType:
		switch t.Dir {
		case ast.SEND:
			identType += "chan<- "
		case ast.RECV:
			identType += "<-chan "
		default:
			identType += "chan "
		}
		identType += tr.resolveType(t.Value)
	case *ast.StarExpr:
		identType += "*"
		identType += tr.resolveType(t.X)
	case *ast.InterfaceType:
		identType += "interface{"
		for idx, method := range t.Methods.List {
			identType += " "
			for nameIdx, name := range method.Names {
				identType += name.String() + " "
				if nameIdx+1 < len(method.Names) {
					identType += ","
				}
			}
			identType += tr.resolveType(method.Type) + " "
			if idx+1 < len(t.Methods.List) {
				identType += ","
			}
		}
		identType += "}"
	default:
		panic(fmt.Sprintf("Not Implemented Type %T", t))
	}
	return identType
}

func getTypeSpec(decl ast.Decl) (ts *ast.TypeSpec, err error) {
	x, ok := decl.(*ast.GenDecl)
	if !ok {
		return nil, fmt.Errorf("unexpected decl type %T", decl)
	}

	if x.Tok != token.TYPE {
		return nil, fmt.Errorf("unexpected identifier %T", x.Tok)
	}
	if len(x.Specs) != 1 {
		return nil, fmt.Errorf("expected one spec, but got %d", len(x.Specs))
	}

	ts, ok = x.Specs[0].(*ast.TypeSpec)
	if !ok {
		return nil, fmt.Errorf("expected type spec, but got %T", x.Specs[0])
	}
	if ts.Name == nil {
		return nil, fmt.Errorf("expected ts name not to be nil")
	}
	return ts, nil
}
