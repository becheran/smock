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
}) *MockWithTypes[T, B] {
	t.Helper()
	m := &MockWithTypes[T, B]{t: t}
	t.Cleanup(func () {
		errStr := ""
		for _, v := range m.vFoo {
			for _, c := range v.expected {
				if c.expectedCalled >= 0 && c.expectedCalled != c.called {
					errStr += fmt.Sprintf("\nExpected 'Foo' to be called %d times, but was called %d times.", c.expectedCalled, c.called)
				}
			}
		}
		for _, v := range m.vEmpty {
			for _, c := range v.expected {
				if c.expectedCalled >= 0 && c.expectedCalled != c.called {
					errStr += fmt.Sprintf("\nExpected 'Empty' to be called %d times, but was called %d times.", c.expectedCalled, c.called)
				}
			}
		}
		if errStr != "" {
			t.Helper()
			t.Fatalf(errStr)
		}})
	return m
}

type MockWithTypes[T any, B any] struct {
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
	
	vFoo []*struct{validateArgs func(a T, b T) bool; expected []*struct{fun func(a T, b T) (r0 B); expectedCalled int; called int}}
	vEmpty []*struct{validateArgs func() bool; expected []*struct{fun func(); expectedCalled int; called int}}
}

func (_this *MockWithTypes[T, B]) Foo(a T, b T) (r0 B) {
	for _, _check := range _this.vFoo {
		if _check.validateArgs == nil || _check.validateArgs(a, b) {
			for _ctr, _exp := range _check.expected {
				if _exp.expectedCalled <= 0 || _ctr == len(_check.expected) - 1 || _exp.called < _exp.expectedCalled {
					_exp.called++
					return _exp.fun(a, b)
				}
			}
		}
	}
	_this.t.Helper()
	_this.unexpectedCall("Foo", a, b)
	return
}

func (_this *MockWithTypes[T, B]) Empty() {
	for _, _check := range _this.vEmpty {
		if _check.validateArgs == nil || _check.validateArgs() {
			for _ctr, _exp := range _check.expected {
				if _exp.expectedCalled <= 0 || _ctr == len(_check.expected) - 1 || _exp.called < _exp.expectedCalled {
					_exp.called++
					_exp.fun()
					return
				}
			}
		}
	}
	_this.t.Helper()
	_this.unexpectedCall("Empty", )
}

func (_this *MockWithTypes[T, B]) unexpectedCall(method string, args ...any) {
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
func (_this *MockWithTypes[T, B]) WHEN() *MockWithTypesWhen[T, B] {
	return &MockWithTypesWhen[T, B] {
		m: _this,
	}
}

type MockWithTypesWhen[T any, B any] struct {
	m *MockWithTypes[T, B]
}

// Defines the behavior when Foo of the mock is called.
//
// As a default the method is expected to be called once.
// To change this behavior use the Times() method to define how often the function shall be called.
func (_this *MockWithTypesWhen[T, B]) Foo() *MockWithTypesFooExpectWithTimes[T, B] {
	for _, f := range _this.m.vFoo {
		if f.validateArgs == nil {
			_this.m.t.Helper()
			_this.m.t.Fatalf("Unreachable condition. Call to 'Foo' is already captured by previous WHEN statement.")
		}
	}
	var defaultExpected struct {
		fun func(a T, b T) (r0 B)
		expectedCalled int
		called int
	}
	defaultExpected.fun = func(a T, b T) (r0 B) { return }
	defaultExpected.expectedCalled = 1
	
	var validator struct {
		validateArgs func(a T, b T) bool
		expected []*struct {
			fun func(a T, b T) (r0 B)
			expectedCalled int
			called int
		}
	}
	validator.expected = append(validator.expected, &defaultExpected)
	_this.m.vFoo = append(_this.m.vFoo, &validator)
	var _then func() *MockWithTypesFooWhen[T, B]
	_then = func() *MockWithTypesFooWhen[T, B] {
		var _newExpected struct {
			fun func(a T, b T) (r0 B)
			expectedCalled int
			called int
		}
		_newExpected.fun = func(a T, b T) (r0 B) { return }
		_newExpected.expectedCalled = 1
		
		validator.expected = append(validator.expected, &_newExpected)
		return &MockWithTypesFooWhen[T, B] {
			expected: validator.expected,
			then: _then,
			t: _this.m.t,
		}
	}
	
	times := &MockWithTypesTimes[*MockWithTypesFooWhen[T, B]] {
		expectedCalled: &validator.expected[0].expectedCalled,
		then: _then,
		t: _this.m.t,
		MockWithTypesThen: MockWithTypesThen[*MockWithTypesFooWhen[T, B]]{ then: _then, t: _this.m.t},
	}
	return &MockWithTypesFooExpectWithTimes[T, B] {
		MockWithTypesFooExpect: &MockWithTypesFooExpect[T, B] {
			MockWithTypesFooWhen: &MockWithTypesFooWhen[T, B] {
				expected: validator.expected,
				then: _then,
				t: _this.m.t,
			},
			validateArgs: &validator.validateArgs,
			times: times,
		},
		MockWithTypesTimes: times,
	}
}

type MockWithTypesFooExpect[T any, B any] struct {
	*MockWithTypesFooWhen[T, B]
	validateArgs *func(a T, b T) bool
	times *MockWithTypesTimes[*MockWithTypesFooWhen[T, B]]
}

// Expect will filter for given arguments.
// Each argument is matched with a filter function. Only if all arguments match this mocked function will be called.

// Arguments are either evaluated using the function, or ignored and always true if the function is set to nil.
func (_this *MockWithTypesFooExpect[T, B]) Expect(a func(T) bool, b func(T) bool) *MockWithTypesFooWhenWithTimes[T, B] {
	if !(a == nil && b == nil) {
		*_this.validateArgs = func(_a T, _b T) bool {
			return (a == nil || a(_a)) && (b == nil || b(_b))
		}
	}
	return &MockWithTypesFooWhenWithTimes[T, B] {
		MockWithTypesFooWhen: _this.MockWithTypesFooWhen,
		MockWithTypesTimes: _this.times,
	}
}

type MockWithTypesFooExpectWithTimes[T any, B any] struct {
	*MockWithTypesTimes[*MockWithTypesFooWhen[T, B]]
	*MockWithTypesFooExpect[T, B]
}

type MockWithTypesFooWhen[T any, B any] struct {
	expected []*struct {
		fun func(a T, b T) (r0 B)
		expectedCalled int
		called int
	}
	then func() *MockWithTypesFooWhen[T, B]
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
}

type MockWithTypesFooWhenWithTimes[T any, B any] struct {
	*MockWithTypesTimes[*MockWithTypesFooWhen[T, B]]
	*MockWithTypesFooWhen[T, B]
}

// Return the provided values when called
func (_this *MockWithTypesFooWhen[T, B]) Return(r0 B) *MockWithTypesTimes[*MockWithTypesFooWhen[T, B]] {
	_this.expected[len(_this.expected) -1].fun = func(T, T) (B) { return r0 }
	return &MockWithTypesTimes[*MockWithTypesFooWhen[T, B]] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockWithTypesThen: MockWithTypesThen[*MockWithTypesFooWhen[T, B]]{ then: _this.then, t: _this.t},
	}
}

// Do will execute the provided function and return the result when called
func (_this *MockWithTypesFooWhen[T, B]) Do(do func(a T, b T) (r0 B)) *MockWithTypesTimes[*MockWithTypesFooWhen[T, B]] {
	_this.expected[len(_this.expected) -1].fun = do
	return &MockWithTypesTimes[*MockWithTypesFooWhen[T, B]] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockWithTypesThen: MockWithTypesThen[*MockWithTypesFooWhen[T, B]]{ then: _this.then, t: _this.t},
	}
}

// Defines the behavior when Empty of the mock is called.
//
// As a default the method is expected to be called once.
// To change this behavior use the Times() method to define how often the function shall be called.
func (_this *MockWithTypesWhen[T, B]) Empty() *MockWithTypesEmptyWhenWithTimes[T, B] {
	for _, f := range _this.m.vEmpty {
		if f.validateArgs == nil {
			_this.m.t.Helper()
			_this.m.t.Fatalf("Unreachable condition. Call to 'Empty' is already captured by previous WHEN statement.")
		}
	}
	var defaultExpected struct {
		fun func()
		expectedCalled int
		called int
	}
	defaultExpected.fun = func() {}
	defaultExpected.expectedCalled = 1
	
	var validator struct {
		validateArgs func() bool
		expected []*struct {
			fun func()
			expectedCalled int
			called int
		}
	}
	validator.expected = append(validator.expected, &defaultExpected)
	_this.m.vEmpty = append(_this.m.vEmpty, &validator)
	var _then func() *MockWithTypesEmptyWhen[T, B]
	_then = func() *MockWithTypesEmptyWhen[T, B] {
		var _newExpected struct {
			fun func()
			expectedCalled int
			called int
		}
		_newExpected.fun = func() { return }
		_newExpected.expectedCalled = 1
		
		validator.expected = append(validator.expected, &_newExpected)
		return &MockWithTypesEmptyWhen[T, B] {
			expected: validator.expected,
			then: _then,
			t: _this.m.t,
		}
	}
	
	times := &MockWithTypesTimes[*MockWithTypesEmptyWhen[T, B]] {
		expectedCalled: &validator.expected[0].expectedCalled,
		then: _then,
		t: _this.m.t,
		MockWithTypesThen: MockWithTypesThen[*MockWithTypesEmptyWhen[T, B]]{ then: _then, t: _this.m.t},
	}
	return &MockWithTypesEmptyWhenWithTimes[T, B] {
		MockWithTypesEmptyWhen: &MockWithTypesEmptyWhen[T, B] {
			expected: validator.expected,
			then: _then,
			t: _this.m.t,
		},
		MockWithTypesTimes: times,
	}
}

type MockWithTypesEmptyWhen[T any, B any] struct {
	expected []*struct {
		fun func()
		expectedCalled int
		called int
	}
	then func() *MockWithTypesEmptyWhen[T, B]
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
}

type MockWithTypesEmptyWhenWithTimes[T any, B any] struct {
	*MockWithTypesTimes[*MockWithTypesEmptyWhen[T, B]]
	*MockWithTypesEmptyWhen[T, B]
}

// Do will execute the provided function and return the result when called
func (_this *MockWithTypesEmptyWhen[T, B]) Do(do func()) *MockWithTypesTimes[*MockWithTypesEmptyWhen[T, B]] {
	_this.expected[len(_this.expected) -1].fun = do
	return &MockWithTypesTimes[*MockWithTypesEmptyWhen[T, B]] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockWithTypesThen: MockWithTypesThen[*MockWithTypesEmptyWhen[T, B]]{ then: _this.then, t: _this.t},
	}
}

type MockWithTypesThen [T any] struct {
	then func() T
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
}

// Then continue with another action
func (_this *MockWithTypesThen[T]) Then() T {
	_this.t.Helper()
	return _this.then()
}

type MockWithTypesTimes[T any] struct {
	expectedCalled *int
	then func() T
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
	MockWithTypesThen[T]
}

// Times sets how often the mocked function is expected to be called.
// Test will fail if the number of calls do not match with the expected calls value.
func (_this *MockWithTypesTimes[T]) Times(times int) *MockWithTypesThen[T] {
	_this.t.Helper()
	*_this.expectedCalled = times
	retVal := &MockWithTypesThen[T] { t: _this.t, then: _this.then }
	if times <= 0 {
		retVal.then = func() T {
			_this.t.Helper()
			callString := "AnyTimes"
			if *_this.expectedCalled == 0 { callString = "Never" }
			_this.t.Fatalf("Then statement is not reachable. Expected calls of previous statement: %s", callString)
			panic("Unreachable!")
		}
	}
	return retVal
}

// AnyTimes disables the check how often a function was called.
func (_this *MockWithTypesTimes[T]) AnyTimes() {
	*_this.expectedCalled = -1
}

// Never will fail if the function is ever called.
func (_this *MockWithTypesTimes[T]) Never() {
	*_this.expectedCalled = 0
}

