package parse

import (
	"fmt"
	"go/ast"
	"go/token"
)

type Reference struct {
	PackageID string // Empty if part of this package
	Name      string // Name of referenced object
}

type Ident struct {
	Name string
	Type string
}

type Method struct {
	TypeParams []Ident
	Params     []Ident
	Results    []Ident
}

type InterfaceResult struct {
	Name       string
	References []Reference
	Methods    []Method
}

func ParseInterface(fset *token.FileSet, file *ast.File, startLine int) (i InterfaceResult, err error) {
	for _, decl := range file.Decls {
		line := fset.Position(decl.Pos()).Line

		if line <= startLine {
			continue
		}

		x, ok := decl.(*ast.GenDecl)
		if !ok {
			return InterfaceResult{}, fmt.Errorf("unexpected decl type %T", decl)
		}

		if x.Tok != token.TYPE {
			return InterfaceResult{}, fmt.Errorf("unexpected identifier %T", x.Tok)
		}
		if len(x.Specs) != 1 {
			return InterfaceResult{}, fmt.Errorf("expected one spec, but got %d", len(x.Specs))
		}
		ts, ok := x.Specs[0].(*ast.TypeSpec)
		if !ok {
			return InterfaceResult{}, fmt.Errorf("expected type spec, but got %T", x.Specs[0])
		}
		if ts.Name == nil {
			return InterfaceResult{}, fmt.Errorf("expected ts name not to be nil")
		}
		name := ts.Name.Name

		interfaceType, ok := ts.Type.(*ast.InterfaceType)
		if !ok {
			if ref := expToReference(ts.Type); ref != nil {
				return InterfaceResult{Name: name, References: []Reference{*ref}}, nil
			}
			return InterfaceResult{}, fmt.Errorf("unexpected type %T", ts.Type)
		}
		// TODO interfaceType.Incomplete?

		references := []Reference{}
		methods := []Method{}
		if interfaceType.Methods != nil {
			for _, it := range interfaceType.Methods.List {
				if ref := expToReference(it.Type); ref != nil {
					references = append(references, *ref)
				} else {
					switch meth := it.Type.(type) {
					case *ast.FuncType:
						methods = append(methods, Method{
							TypeParams: fieldListToIdent(meth.TypeParams),
							Params:     fieldListToIdent(meth.Params),
							Results:    fieldListToIdent(meth.Results),
						})
					default:
						return InterfaceResult{}, fmt.Errorf("unexpected type expression %T", it.Type)
					}
				}
			}
		}

		return InterfaceResult{
			Name:       name,
			Methods:    methods,
			References: references,
		}, nil
	}

	return InterfaceResult{}, fmt.Errorf("interface not found")
}

func expToReference(exp ast.Expr) *Reference {
	switch meth := exp.(type) {
	case *ast.SelectorExpr:
		packageID := ""
		if xIdent, ok := meth.X.(*ast.Ident); ok {
			packageID = xIdent.String()
		}
		return &Reference{
			PackageID: packageID,
			Name:      meth.Sel.String(),
		}
	case *ast.Ident:
		return &Reference{
			Name: meth.String(),
		}
	}
	return nil
}
func fieldListToIdent(list *ast.FieldList) (res []Ident) {
	if list == nil {
		return
	}
	for _, l := range list.List {
		identType := ""
		switch t := l.Type.(type) {
		case *ast.Ident:
			identType = t.String()

		case *ast.SelectorExpr:
			if xIdent, ok := t.X.(*ast.Ident); ok {
				identType = xIdent.String()
			}
			identType += "." + t.Sel.String()
		default:
			continue
		}

		if len(l.Names) == 0 {
			res = append(res, Ident{Type: identType})
		} else {
			for _, name := range l.Names {
				res = append(res, Ident{Name: name.Name, Type: identType})
			}
		}

	}
	return
}
