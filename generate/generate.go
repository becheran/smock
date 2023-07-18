package generate

import (
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/becheran/smock/logger"
	"github.com/becheran/smock/model"
)

func GenerateMock(res model.InterfaceResult) (mock string, err error) {
	logger.Printf("Start generating mock")

	if err := res.ValidateReadyForGenerate(); err != nil {
		return "", err
	}

	hasTypes := len(res.Types) > 0

	mockedStructName := fmt.Sprintf("%s%s", model.MockPrefix, res.Name)
	mockedStructWithTypeIdentifier := fmt.Sprintf("%s%s", mockedStructName, res.Types.ListIdentifier())
	whenStructName := mockedStructName + "When"
	whenStructNameWithTypeIdentifier := fmt.Sprintf("%s%s", whenStructName, res.Types.ListIdentifier())

	w := newWriter()

	version := "unknown"
	if info, found := debug.ReadBuildInfo(); found {
		version = info.Main.Version
	}
	w.P("// DO NOT EDIT")
	w.P("// Code generated by smock %s", version)
	w.P("")

	w.P("package %s%s", res.PackageName, model.MockPackageSuffix)
	w.P("")

	w.P("import (")
	w.Ident()
	fmtAlreadyImported := false
	for _, i := range res.Imports {
		if hasTypes && i.ImportName() == res.PackageName {
			continue
		}
		if i.ImportName() == "fmt" {
			fmtAlreadyImported = true
		}
		w.P("%s", i)
	}
	if !fmtAlreadyImported {
		w.P(`"fmt"`)
	}
	w.EndIdent()
	w.P(")")
	w.P("")

	// Do not validate when generics are used.
	// It is complicated to retrieve a valid type and assert that one concrete type implements the interface.
	if !hasTypes {
		w.P("// %s must implement interface %s.%s", mockedStructName, res.PackageName, res.Name)
		w.P("var _ %s.%s = &%s{}", res.PackageName, res.Name, mockedStructName)
		w.P("")
	}

	w.P(`func New%s%s(t interface {
	Fatalf(format string, args ...interface{})
	Helper()
}) *%s {`,
		mockedStructName, res.Types.ListTypesWithIdentifiers(), mockedStructWithTypeIdentifier)
	w.Ident()
	w.P("return &%s%s{t: t}", mockedStructName, res.Types.ListIdentifier())
	w.EndIdent()
	w.P("}")
	w.P("")

	w.P("type %s%s struct {", mockedStructName, res.Types.ListTypesWithIdentifiers())
	w.Ident()
	w.P(`t interface {
		Fatalf(format string, args ...interface{})
		Helper()
	}`)
	for _, m := range res.Methods {
		w.P("f%s func%s", m.Name, m.Signature())
	}
	w.EndIdent()
	w.P("}")
	w.P("")

	for _, f := range res.Methods {
		w.P("func (m *%s) %s%s {", mockedStructWithTypeIdentifier, f.Name, f.Signature())
		w.Ident()
		w.P("if m.f%s != nil {", f.Name)
		w.Ident()
		retStm := ""
		if len(f.Results) > 0 {
			retStm = "return "
		}
		w.P("%sm.f%s(%s)", retStm, f.Name, f.Params.IdentString(model.IdentTypeInput))
		w.EndIdent()
		w.P("} else {")
		w.Ident()

		args := `fmt.Sprintf("")`
		if len(f.Params) > 0 {
			format := strings.Repeat("%+v, ", len(f.Params))
			format = format[:len(format)-2]
			args = fmt.Sprintf(`fmt.Sprintf("%s", %s)`, format, f.Params.IdentString(model.IdentTypeInput))
		}
		w.P(`m.unexpectedCall("%s", %s)`, f.Name, args)
		w.P(`return`)
		w.EndIdent()
		w.P("}")
		w.EndIdent()
		w.P("}")
		w.P("")
	}

	w.P("func (m *%s) WHEN() *%s {", mockedStructWithTypeIdentifier, whenStructNameWithTypeIdentifier)
	w.Ident()
	w.P("return &%s{", whenStructNameWithTypeIdentifier)
	w.Ident()
	w.P("m: m,")
	w.EndIdent()
	w.P("}")
	w.EndIdent()
	w.P("}")
	w.P("")

	w.P("func (m *%s) unexpectedCall(method, args string) {", mockedStructWithTypeIdentifier)
	w.Ident()
	w.P("m.t.Helper()")
	w.P("m.t.Fatalf(`Unexpected call to %s.%%s(%%s)`, method, args)", mockedStructName)
	w.EndIdent()
	w.P("}")
	w.P("")

	w.P("type %s%s struct {", whenStructName, res.Types.ListTypesWithIdentifiers())
	w.Ident()
	w.P("m *%s", mockedStructWithTypeIdentifier)
	w.EndIdent()
	w.P("}")
	w.P("")

	for idx, f := range res.Methods {
		funcStruct := fmt.Sprintf("%s%sFunc", mockedStructName, f.Name)
		funcStructWithTypeIdentifier := fmt.Sprintf("%s%s", funcStruct, res.Types.ListIdentifier())

		w.P("func (mh *%s) %s() *%s {", whenStructNameWithTypeIdentifier, f.Name, funcStructWithTypeIdentifier)
		w.Ident()
		w.P("mh.m.f%s = func%s { return }", f.Name, f.Signature())
		w.P("return &%s{m: mh.m}", funcStructWithTypeIdentifier)
		w.EndIdent()
		w.P("}")
		w.P("")

		w.P("type %s%s struct {", funcStruct, res.Types.ListTypesWithIdentifiers())
		w.Ident()
		w.P("m *%s", mockedStructWithTypeIdentifier)
		w.EndIdent()
		w.P("}")
		w.P("")

		if len(f.Results) > 0 {
			w.P("func (f *%s) Return(%s) {", funcStructWithTypeIdentifier, f.Results.IdentWithTypeString(model.IdentTypeResult))
			w.Ident()
			w.P("f.m.f%s = func%s { return %s }", f.Name, f.SignatureWithoutIdentifier(), f.Results.IdentString(model.IdentTypeResult))
			w.EndIdent()
			w.P("}")
			w.P("")
		}

		w.P("func (f *%s) Do(do func%s) {", funcStructWithTypeIdentifier, f.Signature())
		w.Ident()
		w.P("f.m.f%s = do", f.Name)
		w.EndIdent()
		w.P("}")
		if idx < len(res.Methods)-1 {
			w.P("")
		}
	}
	logger.Printf("Finished generating mock")

	return w.String(), nil
}
