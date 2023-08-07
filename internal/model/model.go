package model

import (
	"fmt"
	"go/ast"
	"strings"
)

const (
	MockPackageSuffix = "_mock"
	MockDir           = "mocks"
)

type Reference struct {
	PackageID string // Empty if part of this package
	Name      string // Name of referenced object
}

type Ident struct {
	Name string
	Type string
}

type IdentList []Ident

type IdentType string

const (
	IdentTypeInput  IdentType = "i"
	IdentTypeResult IdentType = "r"
)

func (i IdentList) IdentWithTypeStringAndPrefix(identPrefix IdentType, prefix string) (res string) {
	for idx, ident := range i {
		name := ident.Name
		if name == "" {
			name = fmt.Sprintf("%s%d", identPrefix, idx)
		}
		name = prefix + name
		res += fmt.Sprintf("%s %s", name, ident.Type)
		if idx+1 < len(i) {
			res += ", "
		}
	}
	return
}

func (i IdentList) IdentWithTypeString(identPrefix IdentType) (res string) {
	return i.IdentWithTypeStringAndPrefix(identPrefix, "")
}

func (i IdentList) TypeString(identPrefix IdentType) (res string) {
	for idx, ident := range i {
		res += ident.Type
		if idx+1 < len(i) {
			res += ", "
		}
	}
	return
}

func (i IdentList) IdentString(identPrefix IdentType, resolveLambda bool) (res string) {
	for idx, ident := range i {
		name := ident.Name
		if name == "" {
			name = fmt.Sprintf("%s%d", identPrefix, idx)
		}
		res += name
		if resolveLambda && strings.HasPrefix(ident.Type, "...") {
			res += "..."
		}
		if idx+1 < len(i) {
			res += ", "
		}
	}
	return
}

type Method struct {
	Name       string
	TypeParams IdentList
	Params     IdentList
	Results    IdentList
}

func (m Method) Signature() string {
	args := "()"
	if m.Params != nil {
		args = fmt.Sprintf("(%s)", m.Params.IdentWithTypeString(IdentTypeInput))
	}
	retStr := ""
	if len(m.Results) > 0 {
		retStr = fmt.Sprintf(" (%s)", m.Results.IdentWithTypeString(IdentTypeResult))

	}
	return fmt.Sprintf("%s%s", args, retStr)
}

func (m Method) SignatureWithoutIdentifier() string {
	args := "()"
	if m.Params != nil {
		args = fmt.Sprintf("(%s)", m.Params.TypeString(IdentTypeInput))
	}
	retStr := ""
	if len(m.Results) > 0 {
		retStr = fmt.Sprintf(" (%s)", m.Results.TypeString(IdentTypeResult))

	}
	return fmt.Sprintf("%s%s", args, retStr)
}

type Import struct {
	Name string
	Path string
}

func ImportFromAst(spec *ast.ImportSpec) Import {
	name := ""
	if spec.Name != nil {
		name = strings.Trim(spec.Name.Name, `"`)
	}
	return Import{
		Name: name,
		Path: strings.Trim(spec.Path.Value, `"`),
	}
}

func (i Import) ImportName() string {
	if i.Name != "" {
		return i.Name
	}
	idx := strings.LastIndex(i.Path, "/")
	if idx < 0 {
		return i.Path
	}
	name := i.Path[idx+1:]
	if idx := strings.LastIndex(name, "-"); idx >= 0 {
		return name[idx+1:]
	}
	return name
}

func (i Import) String() string {
	return fmt.Sprintf(`%s "%s"`, i.ImportName(), i.Path)
}

type TypesList []Ident

func (tl TypesList) ListIdentifier() (res string) {
	if len(tl) == 0 {
		return
	}
	res += "["
	for idx, t := range tl {
		res += t.Name
		if idx+1 < len(tl) {
			res += ", "
		}
	}
	res += "]"

	return res
}

func (tl TypesList) ListTypesWithIdentifiers() (res string) {
	if len(tl) == 0 {
		return
	}
	res += "["
	for idx, t := range tl {
		res += fmt.Sprintf("%s %s", t.Name, t.Type)
		if idx+1 < len(tl) {
			res += ", "
		}
	}
	res += "]"

	return res
}

type InterfaceResult struct {
	Name        string
	PackageName string
	Imports     []Import
	Methods     []Method
	Types       TypesList
}

// ValidateReadyForGenerate that all members are set and valid
func (ir InterfaceResult) ValidateReadyForGenerate() error {
	if ir.Name == "" {
		return fmt.Errorf("name must be set")
	}
	if ir.PackageName == "" {
		return fmt.Errorf("packageName must be set")
	}
	if len(ir.Imports) < 1 {
		return fmt.Errorf("expected at least one import")
	}
	for _, i := range ir.Imports {
		if i.Path == "" {
			return fmt.Errorf("import path must be set")
		}
	}
	if len(ir.Methods) < 1 {
		return fmt.Errorf("expected at least one method")
	}
	for _, m := range ir.Methods {
		if m.Name == "" {
			return fmt.Errorf("Method name must be set")
		}
	}
	return nil
}
