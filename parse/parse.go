package parse

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/becheran/smock/model"
)

func ParseInterface(fset *token.FileSet, file *ast.File, startLine int) (i model.InterfaceResult, err error) {
	i.PackageName = file.Name.Name

	for _, decl := range file.Decls {
		line := fset.Position(decl.Pos()).Line

		if line <= startLine {
			continue
		}

		x, ok := decl.(*ast.GenDecl)
		if !ok {
			return model.InterfaceResult{}, fmt.Errorf("unexpected decl type %T", decl)
		}

		if x.Tok != token.TYPE {
			return model.InterfaceResult{}, fmt.Errorf("unexpected identifier %T", x.Tok)
		}
		if len(x.Specs) != 1 {
			return model.InterfaceResult{}, fmt.Errorf("expected one spec, but got %d", len(x.Specs))
		}
		ts, ok := x.Specs[0].(*ast.TypeSpec)
		if !ok {
			return model.InterfaceResult{}, fmt.Errorf("expected type spec, but got %T", x.Specs[0])
		}
		if ts.Name == nil {
			return model.InterfaceResult{}, fmt.Errorf("expected ts name not to be nil")
		}
		i.Name = ts.Name.Name

		interfaceType, ok := ts.Type.(*ast.InterfaceType)
		if !ok {
			if ref := expToReference(ts.Type); ref != nil {
				i.References = []model.Reference{*ref}
				return i, nil
			}
			return model.InterfaceResult{}, fmt.Errorf("unexpected type %T", ts.Type)
		}
		// TODO interfaceType.Incomplete?

		if interfaceType.Methods == nil {
			return model.InterfaceResult{}, fmt.Errorf("unexpected empty interface")
		}

		for _, it := range interfaceType.Methods.List {
			if ref := expToReference(it.Type); ref != nil {
				i.References = append(i.References, *ref)
				continue
			}
			if len(it.Names) != 1 {
				continue
			}
			name := it.Names[0]
			if !name.IsExported() {
				continue
			}
			switch meth := it.Type.(type) {
			case *ast.FuncType:
				// TODO filter imports based on used types
				i.Methods = append(i.Methods, model.Method{
					Name:       name.String(),
					TypeParams: fieldListToIdent(meth.TypeParams, i.PackageName),
					Params:     fieldListToIdent(meth.Params, i.PackageName),
					Results:    fieldListToIdent(meth.Results, i.PackageName),
				})
			default:
				return model.InterfaceResult{}, fmt.Errorf("unexpected type expression %T", it.Type)
			}
		}

		return i, nil
	}

	return model.InterfaceResult{}, fmt.Errorf("interface not found")
}

func expToReference(exp ast.Expr) *model.Reference {
	switch meth := exp.(type) {
	case *ast.SelectorExpr:
		packageID := ""
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

func fieldListToIdent(list *ast.FieldList, packageName string) (res []model.Ident) {
	if list == nil {
		return
	}
	for _, l := range list.List {
		tr := typeResolver{PackageName: packageName}
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
	// TODO generic types
}

func (tr *typeResolver) resolveType(exp ast.Expr) (identType string) {
	if exp == nil {
		return "" // For example in case of nil
	}
	switch t := exp.(type) {
	case *ast.Ident:
		// TODO check if is generic
		identType = t.String()
		if t.IsExported() {
			identType = tr.PackageName + "." + identType
		}
	case *ast.SelectorExpr:
		if xIdent, ok := t.X.(*ast.Ident); ok {
			identType = xIdent.String()
		}
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
		// TODO generic types
		identType += "func("
		for idx, param := range t.Params.List {
			identType += tr.resolveType(param.Type)
			if idx+1 < len(t.Params.List) {
				identType += ","
			}
		}
		identType += ")"
		if len(t.Results.List) > 0 {
			identType += " "
			if len(t.Results.List) > 1 {
				identType += "("
			}
			for idx, param := range t.Results.List {
				identType += tr.resolveType(param.Type)
				if idx+1 < len(t.Params.List) {
					identType += ","
				}
			}
			if len(t.Results.List) > 1 {
				identType += ")"
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
	case *ast.InterfaceType:
		identType += "interface{"
		for idx, method := range t.Methods.List {
			identType += " "
			for nameIdx, name := range method.Names {
				// TODO inheritted interfaces
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
