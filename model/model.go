package model

type Reference struct {
	PackageID string // Empty if part of this package
	Name      string // Name of referenced object
}

type Ident struct {
	Name string
	Type string
}

type Method struct {
	Name       string
	TypeParams []Ident
	Params     []Ident
	Results    []Ident
}

func (m Method) Signature() string {
	// TODO
	return ""
}

type Import struct {
	Name string
	Path string
}

type InterfaceResult struct {
	Imports    []Import
	Name       string
	References []Reference
	Methods    []Method
}

func (ir InterfaceResult) ValidateReadyForGenerate() error {
	// TODO
	return nil
}
