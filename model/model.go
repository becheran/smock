package model

import (
	"fmt"
	"strings"
)

const (
	MockPrefix    = "Mock"
	PackageSuffix = "_mock"
	MockDir       = "mocks"
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

func (i IdentList) IdentWithTypeString() (res string) {
	for idx, ident := range i {
		name := ident.Name
		if name == "" {
			name = fmt.Sprintf("r%d", idx)
		}
		res += fmt.Sprintf("%s %s", name, ident.Type)
		if idx+1 < len(i) {
			res += ","
		}
	}
	return
}

func (i IdentList) IdentString() (res string) {
	for idx, ident := range i {
		name := ident.Name
		if name == "" {
			name = fmt.Sprintf("r%d", idx)
		}
		res += name
		if idx+1 < len(i) {
			res += ","
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
	// TODO types?
	// TODO TEST
	args := fmt.Sprintf("(%s)", m.Params.IdentWithTypeString())
	retStr := ""
	if len(m.Results) > 0 {
		retStr = fmt.Sprintf(" (%s)", m.Results.IdentWithTypeString())

	}
	return fmt.Sprintf("%s%s", args, retStr)
}

type Import struct {
	Name string
	Path string
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

type InterfaceResult struct {
	Name        string
	PackageName string
	Imports     []Import
	References  []Reference
	Methods     []Method
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
