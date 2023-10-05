// DO NOT EDIT
// Code generated by smock 

package testpackage_mock

import (
	fmt "fmt"
	io "io"
	os "os"
	reflect "reflect"
	testpackage "github.com/test/testpackage"
)

// MockInheritOwn must implement interface testpackage.InheritOwn
var _ testpackage.InheritOwn = &MockInheritOwn{}

// NewMockInheritOwn creates a new mock object which implements the corresponding interface.
// All function calls can be mocked with a custom behavior for tests using the WHEN function on the mock object.   
func NewMockInheritOwn(t interface {
	Fatalf(format string, args ...any)
	Helper()
	Cleanup(f func())
}) *MockInheritOwn {
	t.Helper()
	m := &MockInheritOwn{t: t}
	t.Cleanup(func () {
		errStr := ""
		for _, v := range m.vRetType {
			if v.expectedCalled >= 0 && v.expectedCalled != v.called {
				errStr += fmt.Sprintf("\nExpected 'RetType' to be called %d times, but was called %d times.", v.expectedCalled, v.called)
			}
		}
		for _, v := range m.vUseStdType {
			if v.expectedCalled >= 0 && v.expectedCalled != v.called {
				errStr += fmt.Sprintf("\nExpected 'UseStdType' to be called %d times, but was called %d times.", v.expectedCalled, v.called)
			}
		}
		if errStr != "" {
			t.Helper()
			t.Fatalf(errStr)
		}
	})
	return m
}

type MockInheritOwn struct {
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
	
	vRetType []*struct{fun func() (r0 testpackage.MyType); validateArgs func() bool; expectedCalled int; called int}
	vUseStdType []*struct{fun func(fi os.FileInfo) (r0 io.Reader); validateArgs func(fi os.FileInfo) bool; expectedCalled int; called int}
}

func (_this *MockInheritOwn) RetType() (r0 testpackage.MyType) {
	for _, _check := range _this.vRetType {
		if _check.validateArgs == nil || _check.validateArgs() {
			_check.called++
			return _check.fun()
		}
	}
	_this.t.Helper()
	_this.unexpectedCall("RetType", )
	return
}

func (_this *MockInheritOwn) UseStdType(fi os.FileInfo) (r0 io.Reader) {
	for _, _check := range _this.vUseStdType {
		if _check.validateArgs == nil || _check.validateArgs(fi) {
			_check.called++
			return _check.fun(fi)
		}
	}
	_this.t.Helper()
	_this.unexpectedCall("UseStdType", fi)
	return
}

func (_this *MockInheritOwn) unexpectedCall(method string, args ...any) {
	argsStr := ""
	for idx, arg := range args {
		switch t := reflect.TypeOf(arg); {
		case t.Kind() == reflect.Func:
			argsStr += fmt.Sprintf("%T", t)
		case t.Kind() == reflect.String:
			argsStr += fmt.Sprintf("%q", arg)
		default:
			argsStr += fmt.Sprintf("%+v", arg)
		}
		if idx+1 < len(args) {
			argsStr += ", "
		}
	}
	_this.t.Helper()
	_this.t.Fatalf(`Unexpected call to %s(). If function call is expected add ".WHEN.%s()" to mock.`, method, method)
}

// WHEN is used to set the mock behavior when a specific functions on the object are called.
// Use this to setup your mock for your specific test scenario.
func (_this *MockInheritOwn) WHEN() *MockInheritOwnWhen {
	return &MockInheritOwnWhen{
		m: _this,
	}
}

type MockInheritOwnWhen struct {
	m *MockInheritOwn
}

// Defines the behavior when RetType of the mock is called.
//
// As a default the method can be called any times.
// To change this behavior use the Times() method to define how often the function shall be called.
func (_this *MockInheritOwnWhen) RetType() *MockInheritOwnRetTypeWhen {
	for _, f := range _this.m.vRetType {
		if f.validateArgs == nil {
			_this.m.t.Helper()
			_this.m.t.Fatalf("Unreachable condition. Call to 'RetType' is already captured by previous WHEN statement.")
		}
	}
	var validator struct {
		fun func() (r0 testpackage.MyType)
		validateArgs func() bool
		expectedCalled int
		called int
	}
	validator.fun = func() (r0 testpackage.MyType) { return }
	validator.expectedCalled = -1
	_this.m.vRetType = append(_this.m.vRetType, &validator)
	return &MockInheritOwnRetTypeWhen{fun: &validator.fun, MockInheritOwnTimes: &MockInheritOwnTimes{expectedCalled: &validator.expectedCalled}} 
}

type MockInheritOwnRetTypeWhen struct {
	*MockInheritOwnTimes
	fun *func() (r0 testpackage.MyType)
}

// Return the provided values when called
func (_this *MockInheritOwnRetTypeWhen) Return(r0 testpackage.MyType) *MockInheritOwnTimes {
	*_this.fun = func() (testpackage.MyType) { return r0 }
	return _this.MockInheritOwnTimes
}

// Do will execute the provided function and return the result when called
func (_this *MockInheritOwnRetTypeWhen) Do(do func() (r0 testpackage.MyType)) *MockInheritOwnTimes {
	*_this.fun = do
	return _this.MockInheritOwnTimes
}

// Defines the behavior when UseStdType of the mock is called.
//
// As a default the method can be called any times.
// To change this behavior use the Times() method to define how often the function shall be called.
func (_this *MockInheritOwnWhen) UseStdType() *MockInheritOwnUseStdTypeExpect {
	for _, f := range _this.m.vUseStdType {
		if f.validateArgs == nil {
			_this.m.t.Helper()
			_this.m.t.Fatalf("Unreachable condition. Call to 'UseStdType' is already captured by previous WHEN statement.")
		}
	}
	var validator struct {
		fun func(fi os.FileInfo) (r0 io.Reader)
		validateArgs func(fi os.FileInfo) bool
		expectedCalled int
		called int
	}
	validator.fun = func(fi os.FileInfo) (r0 io.Reader) { return }
	validator.expectedCalled = -1
	_this.m.vUseStdType = append(_this.m.vUseStdType, &validator)
	return &MockInheritOwnUseStdTypeExpect {
		MockInheritOwnUseStdTypeWhen: &MockInheritOwnUseStdTypeWhen{fun: &validator.fun, MockInheritOwnTimes: &MockInheritOwnTimes{expectedCalled: &validator.expectedCalled}},
		validateArgs: &validator.validateArgs,
	}
}

type MockInheritOwnUseStdTypeExpect struct {
	*MockInheritOwnUseStdTypeWhen
	validateArgs *func(fi os.FileInfo) bool
}

// Expect will filter for given arguments.
// Each argument is matched with a filter function. Only if all arguments match this mocked function will be called.

// Arguments are either evaluated using the function, or ignored and always true if the function is set to nil.
func (_this *MockInheritOwnUseStdTypeExpect) Expect(fi func(os.FileInfo) bool) *MockInheritOwnUseStdTypeWhen {
	if !(fi == nil) {
		*_this.validateArgs = func(_fi os.FileInfo) bool {
			return (fi == nil || fi(_fi))
		}
	}
	return _this.MockInheritOwnUseStdTypeWhen
}

type MockInheritOwnUseStdTypeWhen struct {
	*MockInheritOwnTimes
	fun *func(fi os.FileInfo) (r0 io.Reader)
}

// Return the provided values when called
func (_this *MockInheritOwnUseStdTypeWhen) Return(r0 io.Reader) *MockInheritOwnTimes {
	*_this.fun = func(os.FileInfo) (io.Reader) { return r0 }
	return _this.MockInheritOwnTimes
}

// Do will execute the provided function and return the result when called
func (_this *MockInheritOwnUseStdTypeWhen) Do(do func(fi os.FileInfo) (r0 io.Reader)) *MockInheritOwnTimes {
	*_this.fun = do
	return _this.MockInheritOwnTimes
}

type MockInheritOwnTimes struct {
	expectedCalled *int
}

// Times sets how often the mocked function is expected to be called.
// Test will fail if the number of calls do not match with the expected calls value.
//
// A number < 0 means that a function may be called any times which is also the default behavior.
func (_this *MockInheritOwnTimes) Times(times int) {
	*_this.expectedCalled = times
}

// AnyTimes disables the check how often a function was called.
func (_this *MockInheritOwnTimes) AnyTimes() {
	*_this.expectedCalled = -1
}

// Never will fail if the function is ever called. Is the same as Times(0).
func (_this *MockInheritOwnTimes) Never() {
	*_this.expectedCalled = 0
}

// Once will fail if the function is not called once. Is the same as Times(1).
func (_this *MockInheritOwnTimes) Once() {
	*_this.expectedCalled = 1
}
