package generate

import (
	"fmt"
	"go/token"
	"runtime/debug"
	"sort"
	"strings"

	"github.com/becheran/smock/internal/logger"
	"github.com/becheran/smock/internal/model"
)

func GenerateMock(res model.InterfaceResult) (mock []byte, err error) {
	logger.Printf("Start generating mock for %+v", res)

	if err := res.ValidateReadyForGenerate(); err != nil {
		return nil, err
	}

	hasTypes := len(res.Types) > 0
	assertImplements := !(hasTypes || !token.IsExported(res.Name))

	mockedStructName := fmt.Sprintf("Mock%s", res.Name)
	mockedStructWithTypeIdentifier := fmt.Sprintf("%s%s", mockedStructName, res.Types.ListIdentifier())
	whenStructName := mockedStructName + "When"
	whenStructNameWithTypeIdentifier := fmt.Sprintf("%s%s", whenStructName, res.Types.ListIdentifier())

	w := newWriter()

	printHelperInterface := func() {
		w.P("t interface {")
		w.Ident()
		w.P("Fatalf(format string, args ...any)")
		w.P("Helper()")
		w.EndIdent()
		w.P("}")
	}

	version := "unknown"
	if info, found := debug.ReadBuildInfo(); found {
		version = info.Main.Version
	}
	w.P("// DO NOT EDIT")
	w.P("// Code generated by smock %s", version)
	w.P("")

	w.P("package %s%s", res.PackageName, model.MockPackageSuffix)
	w.P("")

	fmtAlreadyImported := false
	reflectAlreadyImported := false
	for _, i := range res.Imports {
		if !assertImplements && i.ImportName() == res.PackageName {
			continue
		}
		if i.ImportName() == "fmt" {
			fmtAlreadyImported = true
		}
		if i.ImportName() == "reflect" {
			reflectAlreadyImported = true
		}
	}
	if !fmtAlreadyImported {
		res.Imports = append(res.Imports, model.Import{Path: "fmt"})
	}
	if !reflectAlreadyImported {
		res.Imports = append(res.Imports, model.Import{Path: "reflect"})
	}
	sort.SliceStable(res.Imports, func(a, b int) bool {
		return strings.Compare(res.Imports[a].ImportName(), res.Imports[b].ImportName()) < 0
	})

	w.P("import (")
	w.Ident()
	for _, i := range res.Imports {
		if !assertImplements && i.ImportName() == res.PackageName {
			continue
		}
		w.P("%s", i)
		logger.Printf("Use import: %s", i)
	}
	w.EndIdent()
	w.P(")")
	w.P("")

	// Do not validate when generics are used.
	// It is complicated to retrieve a valid type and assert that one concrete type implements the interface.
	if assertImplements {
		w.P("// %s must implement interface %s.%s", mockedStructName, res.PackageName, res.Name)
		w.P("var _ %s.%s = &%s{}", res.PackageName, res.Name, mockedStructName)
		w.P("")
	}

	w.P(`// NewMock%s creates a new mock object which implements the corresponding interface.
// All function calls can be mocked with a custom behavior for tests using the WHEN function on the mock object.   
func NewMock%s%s(t interface {
	Fatalf(format string, args ...any)
	Helper()
	Cleanup(f func())
}) *%s {`,
		res.Name, res.Name, res.Types.ListTypesWithIdentifiers(), mockedStructWithTypeIdentifier)
	w.Ident()
	w.P("t.Helper()")
	w.P("m := &%s%s{t: t}", mockedStructName, res.Types.ListIdentifier())
	w.P("t.Cleanup(func () {")
	w.Ident()
	w.P("errStr := \"\"")
	for _, m := range res.Methods {
		w.P("for _, v := range m.v%s {", m.Name)
		w.Ident()
		w.P("for _, c := range v.expected {")
		w.Ident()
		w.P("if c.expectedCalled >= 0 && c.expectedCalled != c.called {")
		w.Ident()
		// TODO add args?
		w.P("errStr += fmt.Sprintf(\"\\nExpected '%s' to be called %%d times, but was called %%d times.\", c.expectedCalled, c.called)", m.Name)
		w.EndIdent()
		w.P("}")
		w.EndIdent()
		w.P("}")
		w.EndIdent()
		w.P("}")
	}
	w.P("if errStr != \"\" {")
	w.Ident()
	w.P("t.Helper()")
	w.P("t.Fatalf(errStr)")
	w.EndIdent()
	w.P("}})")
	w.EndIdent()
	w.P("return m")
	w.EndIdent()
	w.P("}")
	w.EndIdent()
	w.P("")

	w.P("type %s%s struct {", mockedStructName, res.Types.ListTypesWithIdentifiers())
	w.Ident()
	printHelperInterface()
	w.P("")
	for _, m := range res.Methods {
		w.P("v%s []*struct{validateArgs func(%s) bool; expected []*struct{fun func%s; expectedCalled int; called int}}",
			m.Name, m.Params.IdentWithTypeString(model.IdentTypeInput), m.Signature())
	}
	w.EndIdent()
	w.P("}")
	w.P("")

	for _, f := range res.Methods {
		w.P("func (_this *%s) %s%s {", mockedStructWithTypeIdentifier, f.Name, f.Signature())
		w.Ident()
		w.P("for _, _check := range _this.v%s {", f.Name)
		w.Ident()
		w.P("if _check.validateArgs == nil || _check.validateArgs(%s) {", f.Params.IdentString(model.IdentTypeInput, true))
		w.Ident()
		w.P("for _ctr, _exp := range _check.expected {")
		w.Ident()
		w.P("if _exp.expectedCalled <= 0 || _ctr == len(_check.expected) - 1 || _exp.called < _exp.expectedCalled {")
		w.Ident()
		w.P("_exp.called++")
		if len(f.Results) > 0 {
			w.P("return _exp.fun(%s)", f.Params.IdentString(model.IdentTypeInput, true))
		} else {
			w.P("_exp.fun(%s)", f.Params.IdentString(model.IdentTypeInput, true))
			w.P("return")
		}
		w.EndIdent()
		w.P("}")
		w.EndIdent()
		w.P("}")
		w.EndIdent()
		w.P("}")
		w.EndIdent()
		w.P("}")
		w.P("_this.t.Helper()")
		w.P(`_this.unexpectedCall("%s", %s)`, f.Name, f.Params.IdentString(model.IdentTypeInput, false))
		if len(f.Results) > 0 {
			w.P(`return`)
		}
		w.EndIdent()
		w.P("}")
		w.P("")
	}

	w.P("func (_this *%s) unexpectedCall(method string, args ...any) {", mockedStructWithTypeIdentifier)
	w.Ident()
	w.P("argsStr := \"\"")
	w.P("for idx, arg := range args {")
	w.Ident()
	w.P("switch t := reflect.TypeOf(arg); {")
	w.P("case t.Kind() == reflect.Func:")
	w.Ident()
	w.P("argsStr += fmt.Sprintf(\"%%T\", t)")
	w.EndIdent()
	w.P("case t.Kind() == reflect.String:")
	w.Ident()
	w.P("argsStr += fmt.Sprintf(\"%%q\", arg)")
	w.EndIdent()
	w.P("default:")
	w.Ident()
	w.P("argsStr += fmt.Sprintf(\"%%+v\", arg)")
	w.EndIdent()
	w.P("}")
	w.P("if idx+1 < len(args) {")
	w.Ident()
	w.P("argsStr += \", \"")
	w.EndIdent()
	w.P("}")
	w.EndIdent()
	w.P("}")
	w.P("_this.t.Helper()")
	w.P("_this.t.Fatalf(`Unexpected call to %%s(). If function call is expected add \".WHEN.%%s()\" to mock.`, method, method)")
	w.EndIdent()
	w.P("}")
	w.P("")

	w.P("// WHEN is used to set the mock behavior when a specific functions on the object are called.")
	w.P("// Use this to setup your mock for your specific test scenario.")
	w.P("func (_this *%s) WHEN() *%s {", mockedStructWithTypeIdentifier, whenStructNameWithTypeIdentifier)
	w.Ident()
	w.P("return &%s {", whenStructNameWithTypeIdentifier)
	w.Ident()
	w.P("m: _this,")
	w.EndIdent()
	w.P("}")
	w.EndIdent()
	w.P("}")
	w.P("")

	w.P("type %s%s struct {", whenStructName, res.Types.ListTypesWithIdentifiers())
	w.Ident()
	w.P("m *%s", mockedStructWithTypeIdentifier)
	w.EndIdent()
	w.P("}")
	w.P("")

	timesStruct := fmt.Sprintf("%s%s", mockedStructName, "Times")
	thenStruct := fmt.Sprintf("%sThen", mockedStructName)

	for idx, f := range res.Methods {
		expectStruct := fmt.Sprintf("%s%sExpect", mockedStructName, f.Name)
		expectStructRef := fmt.Sprintf("%s%s", expectStruct, res.Types.ListIdentifier())
		expectStructWithTimes := fmt.Sprintf("%sWithTimes", expectStruct)
		expectStructWithTimesRef := fmt.Sprintf("%s%s", expectStructWithTimes, res.Types.ListIdentifier())
		whenStruct := fmt.Sprintf("%s%sWhen", mockedStructName, f.Name)
		whenStructRef := fmt.Sprintf("%s%s", whenStruct, res.Types.ListIdentifier())
		whenStructWithTimes := fmt.Sprintf("%sWithTimes", whenStruct)
		whenStructWithTimesRef := fmt.Sprintf("%s%s", whenStructWithTimes, res.Types.ListIdentifier())

		hasParams := len(f.Params) > 0
		hasReturnValues := len(f.Results) > 0

		ref := whenStructWithTimesRef
		if hasParams {
			ref = expectStructWithTimesRef
		}
		w.P("// Defines the behavior when %s of the mock is called.", f.Name)
		w.P("//")
		w.P("// As a default the method is expected to be called once.")
		w.P("// To change this behavior use the `Times()` method to define how often the function shall be called.")
		w.P("func (_this *%s) %s() *%s {", whenStructNameWithTypeIdentifier, f.Name, ref)
		w.Ident()
		w.P("for _, f := range _this.m.v%s {", f.Name)
		w.Ident()
		w.P("if f.validateArgs == nil {")
		w.Ident()
		w.P("_this.m.t.Helper()")
		w.P("_this.m.t.Fatalf(\"Unreachable condition. Call to '%s' is already captured by previous WHEN statement.\")", f.Name)
		w.EndIdent()
		w.P("}")
		w.EndIdent()
		w.P("}")

		w.P("var defaultExpected struct {")
		w.Ident()
		w.P("fun func%s", f.Signature())
		w.P("expectedCalled int")
		w.P("called int")
		w.EndIdent()
		w.P("}")
		ret := ""
		if hasReturnValues {
			ret = " return "
		}
		w.P("defaultExpected.fun = func%s {%s}", f.Signature(), ret)
		w.P("defaultExpected.expectedCalled = 1")
		w.P("")
		w.P("var validator struct {")
		w.Ident()
		w.P("validateArgs func(%s) bool", f.Params.IdentWithTypeString(model.IdentTypeInput))
		w.P("expected []*struct {")
		w.Ident()
		w.P("fun func%s", f.Signature())
		w.P("expectedCalled int")
		w.P("called int")
		w.EndIdent()
		w.P("}")
		w.EndIdent()
		w.P("}")
		w.P("validator.expected = append(validator.expected, &defaultExpected)")
		w.P("_this.m.v%s = append(_this.m.v%s, &validator)", f.Name, f.Name)
		actionStruct := whenStruct
		if hasParams {
			actionStruct = expectStruct
		}
		actionStructRef := fmt.Sprintf("%s%s", actionStruct, res.Types.ListIdentifier())
		w.P("var _then func() *%s", whenStructRef)
		w.P("_then = func() *%s {", whenStructRef)
		w.Ident()
		w.P("var _newExpected struct {")
		w.Ident()
		w.P("fun func%s", f.Signature())
		w.P("expectedCalled int")
		w.P("called int")
		w.EndIdent()
		w.P("}")
		w.P("_newExpected.fun = func%s { return }", f.Signature())
		w.P("_newExpected.expectedCalled = 1")
		w.P("")
		w.P("validator.expected = append(validator.expected, &_newExpected)")
		w.P("return &%s {", whenStructRef)
		w.Ident()
		w.P("expected: validator.expected,")
		w.P("then: _then,")
		w.P("t: _this.m.t,")
		w.EndIdent()
		w.P("}")
		w.EndIdent()
		w.P("}")
		w.P("")

		returnType := whenStructWithTimesRef
		if hasParams {
			returnType = expectStructWithTimesRef
		}

		w.P("times := &%s[*%s] {", timesStruct, whenStructRef)
		w.Ident()
		w.P("expectedCalled: &validator.expected[0].expectedCalled,")
		w.P("then: _then,")
		w.P("t: _this.m.t,")
		w.P("%s: %s[*%s]{ then: _then, t: _this.m.t},", thenStruct, thenStruct, whenStructRef)
		w.EndIdent()
		w.P("}")

		w.P("return &%s {", returnType)
		w.Ident()
		w.P("%s: &%s {", actionStruct, actionStructRef)
		w.Ident()
		if hasParams {
			w.P("%s: &%s {", whenStruct, whenStructRef)
			w.Ident()
			w.P("expected: validator.expected,")
			w.P("then: _then,")
			w.P("t: _this.m.t,")
			w.EndIdent()
			w.P("},")
			w.P("validateArgs: &validator.validateArgs,")
			w.P("times: times,")
		} else {
			w.P("expected: validator.expected,")
			w.P("then: _then,")
			w.P("t: _this.m.t,")
		}
		w.EndIdent()
		w.P("},")
		w.P("%s: times,", timesStruct)
		w.EndIdent()
		w.P("}")
		w.EndIdent()
		w.P("}")
		w.P("")

		if hasParams {
			w.P("type %s%s struct {", expectStruct, res.Types.ListTypesWithIdentifiers())
			w.Ident()
			w.P("*%s", whenStructRef)
			w.P("validateArgs *func(%s) bool", f.Params.IdentWithTypeString(model.IdentTypeInput))
			w.P("times *%s[*%s]", timesStruct, whenStructRef)
			w.EndIdent()
			w.P("}")
			w.P("")

			args := ""
			lambdaFieldName := ""
			for idx, arg := range f.Params {
				name := arg.Name
				typeStr := arg.Type
				lambdaPref := ""
				if strings.HasPrefix(arg.Type, "...") {
					lambdaFieldName = name
					typeStr = strings.TrimPrefix(typeStr, "...")
					lambdaPref = "..."
				}
				if name == "" {
					name = fmt.Sprintf("_%d", idx)
				}
				args += fmt.Sprintf("%s %sfunc(%s) bool", name, lambdaPref, typeStr)
				if idx+1 < len(f.Params) {
					args += ", "
				}
			}
			w.P("// Expect will filter for given arguments.")
			w.P("// Each argument is matched with a filter function. Only if all arguments match this mocked function will be called.")
			w.P("")
			w.P("// Arguments are either evaluated using the function, or ignored and always true if the function is set to nil.")
			w.P("func (_this *%s) Expect(%s) *%s {", expectStructRef, args, whenStructWithTimesRef)
			w.Ident()
			matchString := ""
			checkAllNil := ""
			for idx, arg := range f.Params {
				if strings.HasPrefix(arg.Type, "...") {
					checkAllNil += fmt.Sprintf("len(%s) == 0", lambdaFieldName)
					matchString += "true"
					break
				}
				name := arg.Name
				input := arg.Name
				if name == "" {
					name = fmt.Sprintf("_%d", idx)
					input = fmt.Sprintf("i%d", idx)
				}
				if strings.HasPrefix(arg.Type, "...") {
					input += "..."
				}
				matchString += fmt.Sprintf("(%s == nil || %s(_%s))", name, name, input)
				checkAllNil += fmt.Sprintf("%s == nil", name)
				if idx+1 < len(f.Params) {
					matchString += " && "
					checkAllNil += " && "
				}
			}
			w.P("if !(%s) {", checkAllNil)
			w.Ident()
			w.P("*_this.validateArgs = func(%s) bool {", f.Params.IdentWithTypeStringAndPrefix(model.IdentTypeInput, "_"))
			w.Ident()
			if lambdaFieldName != "" {
				w.P("for _idx, _val := range _%s {", lambdaFieldName)
				w.Ident()
				w.P("if _idx >= len(%s) || !(%s[_idx] == nil || %s[_idx](_val)) {", lambdaFieldName, lambdaFieldName, lambdaFieldName)
				w.Ident()
				w.P("return false")
				w.EndIdent()
				w.P("}")
				w.EndIdent()
				w.P("}")
			}
			w.P("return %s", matchString)
			w.EndIdent()
			w.P("}")
			w.EndIdent()
			w.P("}")
			w.P("return &%s {", whenStructWithTimesRef)
			w.Ident()
			w.P("%s: _this.%s,", whenStruct, whenStruct)
			w.P("%s: _this.times,", timesStruct)
			w.EndIdent()
			w.P("}")
			w.EndIdent()
			w.P("}")
			w.P("")

			w.P("type %s%s struct {", expectStructWithTimes, res.Types.ListTypesWithIdentifiers())
			w.Ident()
			w.P("*%s[*%s]", timesStruct, whenStructRef)
			w.P("*%s", expectStructRef)
			w.EndIdent()
			w.P("}")
			w.P("")
		}

		w.P("type %s%s struct {", whenStruct, res.Types.ListTypesWithIdentifiers())
		w.Ident()
		w.P("expected []*struct {")
		w.Ident()
		w.P("fun func%s", f.Signature())
		w.P("expectedCalled int")
		w.P("called int")
		w.EndIdent()
		w.P("}")
		w.P("then func() *%s", whenStructRef)
		printHelperInterface()
		w.EndIdent()
		w.P("}")
		w.P("")

		w.P("type %s%s struct {", whenStructWithTimes, res.Types.ListTypesWithIdentifiers())
		w.Ident()
		w.P("*%s[*%s]", timesStruct, whenStructRef)
		w.P("*%s", whenStructRef)
		w.EndIdent()
		w.P("}")
		w.P("")

		printReturnFromFunc := func() {
			w.P("return &%s[*%s] {", timesStruct, whenStructRef)
			w.Ident()
			w.P("expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,")
			w.P("then: _this.then,")
			w.P("t: _this.t,")
			w.P("%s: %s[*%s]{ then: _this.then, t: _this.t},", thenStruct, thenStruct, whenStructRef)
			w.EndIdent()
			w.P("}")
		}

		if hasReturnValues {
			w.P("// Return the provided values when called")
			w.P("func (_this *%s) Return(%s) *%s[*%s] {", whenStructRef, f.Results.IdentWithTypeString(model.IdentTypeResult), timesStruct, whenStructRef)
			w.Ident()
			w.P("_this.expected[len(_this.expected) -1].fun = func%s { return %s }", f.SignatureWithoutIdentifier(), f.Results.IdentString(model.IdentTypeResult, false))
			printReturnFromFunc()
			w.EndIdent()
			w.P("}")
			w.P("")
		}

		w.P("// Do will execute the provided function and return the result when called")
		w.P("func (_this *%s) Do(do func%s) *%s[*%s] {", whenStructRef, f.Signature(), timesStruct, whenStructRef)
		w.Ident()
		w.P("_this.expected[len(_this.expected) -1].fun = do")
		printReturnFromFunc()
		w.EndIdent()
		w.P("}")
		if idx < len(res.Methods)-1 {
			w.P("")
		}
	}

	w.P("")
	w.P("type %s [T any] struct {", thenStruct)
	w.Ident()
	w.P("then func() T")
	printHelperInterface()
	w.EndIdent()
	w.P("}")
	w.P("")
	w.P("// Then continue with another action")
	w.P("func (_this *%s[T]) Then() T {", thenStruct)
	w.Ident()
	w.P("_this.t.Helper()")
	w.P("return _this.then()")
	w.EndIdent()
	w.P("}")

	w.P("")
	w.P("type %s[T any] struct {", timesStruct)
	w.Ident()
	w.P("expectedCalled *int")
	w.P("then func() T")
	printHelperInterface()
	w.P("%s[T]", thenStruct)
	w.EndIdent()
	w.P("}")
	w.P("")

	w.P("// Times sets how often the mocked function is expected to be called.")
	w.P("// Test will fail if the number of calls do not match with the expected calls value.")
	w.P("func (_this *%s[T]) Times(times int) *%s[T] {", timesStruct, thenStruct)
	w.Ident()
	w.P("_this.t.Helper()")
	w.P("*_this.expectedCalled = times")
	w.P("retVal := &%s[T] { t: _this.t, then: _this.then }", thenStruct)
	w.P("if times <= 0 {")
	w.Ident()
	w.P("retVal.then = func() T {")
	w.Ident()
	w.P("_this.t.Helper()")
	w.P(`callString := "AnyTimes"`)
	w.P(`if *_this.expectedCalled == 0 { callString = "Never" }`)
	w.P(`_this.t.Fatalf("Then statement is not reachable. Expected calls of previous statement: %%s", callString)`)
	w.P(`panic("Unreachable!")`)
	w.EndIdent()
	w.P("}")
	w.EndIdent()
	w.P("}")
	w.P("return retVal")
	w.EndIdent()
	w.P("}")
	w.P("")

	w.P("// AnyTimes disables the check how often a function was called.")
	w.P("func (_this *%s[T]) AnyTimes() {", timesStruct)
	w.Ident()
	w.P("*_this.expectedCalled = -1")
	w.EndIdent()
	w.P("}")
	w.P("")

	w.P("// Never will fail if the function is ever called.")
	w.P("func (_this *%s[T]) Never() {", timesStruct)
	w.Ident()
	w.P("*_this.expectedCalled = 0")
	w.EndIdent()
	w.P("}")
	w.P("")

	logger.Printf("Finished generating mock")

	return w.buff.Bytes(), nil
}
