// DO NOT EDIT
// Code generated by smock 

package testpackage_mock

import (
	fmt "fmt"
	reflect "reflect"
)

// NewMockWithTypes creates a new mock object which implements the corresponding interface.
// All function calls can be mocked with a custom behavior for tests using the WHEN function on the mock object.   
func NewMockWithTypes[T any, B any](t interface {
	Fatalf(format string, args ...any)
	Helper()
	Cleanup(f func())
}) *mockWithTypes[T, B] {
	t.Helper()
	m := &mockWithTypes[T, B]{t: t}
	t.Cleanup(func () {
		errStr := ""
		for _, v := range m.vFoo {
			if v.expectedCalled >= 0 && v.expectedCalled != v.called {
				errStr += fmt.Sprintf("\nExpected 'Foo' to be called %d times, but was called %d times.", v.expectedCalled, v.called)
			}
		}
		for _, v := range m.vEmpty {
			if v.expectedCalled >= 0 && v.expectedCalled != v.called {
				errStr += fmt.Sprintf("\nExpected 'Empty' to be called %d times, but was called %d times.", v.expectedCalled, v.called)
			}
		}
		if errStr != "" {
			t.Helper()
			t.Fatalf(errStr)
		}
	})
	return m
}

type mockWithTypes[T any, B any] struct {
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
	
	vFoo []*struct{fun func(a T, b T) (r0 B); validateArgs func(a T, b T) bool; expectedCalled int; called int}
	vEmpty []*struct{fun func(); validateArgs func() bool; expectedCalled int; called int}
}

func (_this *mockWithTypes[T, B]) Foo(a T, b T) (r0 B) {
	for _, _check := range _this.vFoo {
		if _check.validateArgs == nil || _check.validateArgs(a, b) {
			_check.called++
			return _check.fun(a, b)
		}
	}
	_this.t.Helper()
	_this.unexpectedCall("Foo", a, b)
	return
}

func (_this *mockWithTypes[T, B]) Empty() {
	for _, _check := range _this.vEmpty {
		if _check.validateArgs == nil || _check.validateArgs() {
			_check.called++
			_check.fun()
			return
		}
	}
	_this.t.Helper()
	_this.unexpectedCall("Empty", )
}

func (_this *mockWithTypes[T, B]) unexpectedCall(method string, args ...any) {
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
func (_this *mockWithTypes[T, B]) WHEN() *mockWithTypesWhen[T, B] {
	return &mockWithTypesWhen[T, B]{
		m: _this,
	}
}

type mockWithTypesWhen[T any, B any] struct {
	m *mockWithTypes[T, B]
}

// Defines the behavior when Foo of the mock is called.
//
// As a default the method can be called any times.
// To change this behavior use the Times() method to define how often the function shall be called.
func (_this *mockWithTypesWhen[T, B]) Foo() *mockWithTypesFooExpect[T, B] {
	for _, f := range _this.m.vFoo {
		if f.validateArgs == nil {
			_this.m.t.Helper()
			_this.m.t.Fatalf("Unreachable condition. Call to 'Foo' is already captured by previous WHEN statement.")
		}
	}
	var validator struct {
		fun func(a T, b T) (r0 B)
		validateArgs func(a T, b T) bool
		expectedCalled int
		called int
	}
	validator.fun = func(a T, b T) (r0 B) { return }
	validator.expectedCalled = -1
	_this.m.vFoo = append(_this.m.vFoo, &validator)
	return &mockWithTypesFooExpect[T, B] {
		mockWithTypesFooWhen: &mockWithTypesFooWhen[T, B]{fun: &validator.fun, mockWithTypesTimes: &mockWithTypesTimes{expectedCalled: &validator.expectedCalled}},
		validateArgs: &validator.validateArgs,
	}
}

type mockWithTypesFooExpect[T any, B any] struct {
	*mockWithTypesFooWhen[T, B]
	validateArgs *func(a T, b T) bool
}

// Expect will filter for given arguments.
// Each argument is matched with a filter function. Only if all arguments match this mocked function will be called.

// Arguments are either evaluated using the function, or ignored and always true if the function is set to nil.
func (_this *mockWithTypesFooExpect[T, B]) Expect(a func(T) bool, b func(T) bool) *mockWithTypesFooWhen[T, B] {
	if !(a == nil && b == nil) {
		*_this.validateArgs = func(_a T, _b T) bool {
			return (a == nil || a(_a)) && (b == nil || b(_b))
		}
	}
	return _this.mockWithTypesFooWhen
}

type mockWithTypesFooWhen[T any, B any] struct {
	*mockWithTypesTimes
	fun *func(a T, b T) (r0 B)
}

// Return the provided values when called
func (_this *mockWithTypesFooWhen[T, B]) Return(r0 B) *mockWithTypesTimes {
	*_this.fun = func(T, T) (B) { return r0 }
	return _this.mockWithTypesTimes
}

// Do will execute the provided function and return the result when called
func (_this *mockWithTypesFooWhen[T, B]) Do(do func(a T, b T) (r0 B)) *mockWithTypesTimes {
	*_this.fun = do
	return _this.mockWithTypesTimes
}

// Defines the behavior when Empty of the mock is called.
//
// As a default the method can be called any times.
// To change this behavior use the Times() method to define how often the function shall be called.
func (_this *mockWithTypesWhen[T, B]) Empty() *mockWithTypesEmptyWhen[T, B] {
	for _, f := range _this.m.vEmpty {
		if f.validateArgs == nil {
			_this.m.t.Helper()
			_this.m.t.Fatalf("Unreachable condition. Call to 'Empty' is already captured by previous WHEN statement.")
		}
	}
	var validator struct {
		fun func()
		validateArgs func() bool
		expectedCalled int
		called int
	}
	validator.fun = func() {}
	validator.expectedCalled = -1
	_this.m.vEmpty = append(_this.m.vEmpty, &validator)
	return &mockWithTypesEmptyWhen[T, B]{fun: &validator.fun, mockWithTypesTimes: &mockWithTypesTimes{expectedCalled: &validator.expectedCalled}} 
}

type mockWithTypesEmptyWhen[T any, B any] struct {
	*mockWithTypesTimes
	fun *func()
}

// Do will execute the provided function and return the result when called
func (_this *mockWithTypesEmptyWhen[T, B]) Do(do func()) *mockWithTypesTimes {
	*_this.fun = do
	return _this.mockWithTypesTimes
}

type mockWithTypesTimes struct {
	expectedCalled *int
}

// Times sets how often the mocked function is expected to be called.
// Test will fail if the number of calls do not match with the expected calls value.
//
// A number < 0 means that a function may be called any times which is also the default behavior.
func (_this *mockWithTypesTimes) Times(times int) {
	*_this.expectedCalled = times
}

// AnyTimes disables the check how often a function was called.
func (_this *mockWithTypesTimes) AnyTimes() {
	*_this.expectedCalled = -1
}

// Never will fail if the function is ever called. Is the same as Times(0).
func (_this *mockWithTypesTimes) Never() {
	*_this.expectedCalled = 0
}

// Once will fail if the function is not called once. Is the same as Times(1).
func (_this *mockWithTypesTimes) Once() {
	*_this.expectedCalled = 1
}
