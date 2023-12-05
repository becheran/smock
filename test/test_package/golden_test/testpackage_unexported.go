// DO NOT EDIT
// Code generated by smock 

package testpackage_mock

import (
	fmt "fmt"
	reflect "reflect"
)

// NewMockunexported creates a new mock object which implements the corresponding interface.
// All function calls can be mocked with a custom behavior for tests using the WHEN function on the mock object.   
func NewMockunexported(t interface {
	Fatalf(format string, args ...any)
	Helper()
	Cleanup(f func())
}) *Mockunexported {
	t.Helper()
	m := &Mockunexported{t: t}
	t.Cleanup(func () {
		errStr := ""
		for _, v := range m.vFoo {
			for _, c := range v.expected {
				if c.expectedCalled >= 0 && c.expectedCalled != c.called {
					errStr += fmt.Sprintf("\nExpected 'Foo' to be called %d times, but was called %d times.", c.expectedCalled, c.called)
				}
			}
		}
		if errStr != "" {
			t.Helper()
			t.Fatalf(errStr)
		}})
	return m
}

type Mockunexported struct {
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
	
	vFoo []*struct{validateArgs func() bool; expected []*struct{fun func(); expectedCalled int; called int}}
}

func (_this *Mockunexported) Foo() {
	for _, _check := range _this.vFoo {
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
	_this.unexpectedCall("Foo", )
}

func (_this *Mockunexported) unexpectedCall(method string, args ...any) {
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
func (_this *Mockunexported) WHEN() *MockunexportedWhen {
	return &MockunexportedWhen {
		m: _this,
	}
}

type MockunexportedWhen struct {
	m *Mockunexported
}

// Defines the behavior when Foo of the mock is called.
//
// As a default the method is expected to be called once.
// To change this behavior use the `Times()` method to define how often the function shall be called.
func (_this *MockunexportedWhen) Foo() *MockunexportedFooWhenWithTimes {
	for _, f := range _this.m.vFoo {
		if f.validateArgs == nil {
			_this.m.t.Helper()
			_this.m.t.Fatalf("Unreachable condition. Call to 'Foo' is already captured by previous WHEN statement.")
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
	_this.m.vFoo = append(_this.m.vFoo, &validator)
	var _then func() *MockunexportedFooWhen
	_then = func() *MockunexportedFooWhen {
		var _newExpected struct {
			fun func()
			expectedCalled int
			called int
		}
		_newExpected.fun = func() { return }
		_newExpected.expectedCalled = 1
		
		validator.expected = append(validator.expected, &_newExpected)
		return &MockunexportedFooWhen {
			expected: validator.expected,
			then: _then,
			t: _this.m.t,
		}
	}
	
	times := &MockunexportedTimes[*MockunexportedFooWhen] {
		expectedCalled: &validator.expected[0].expectedCalled,
		then: _then,
		t: _this.m.t,
		MockunexportedThen: MockunexportedThen[*MockunexportedFooWhen]{ then: _then, t: _this.m.t},
	}
	return &MockunexportedFooWhenWithTimes {
		MockunexportedFooWhen: &MockunexportedFooWhen {
			expected: validator.expected,
			then: _then,
			t: _this.m.t,
		},
		MockunexportedTimes: times,
	}
}

type MockunexportedFooWhen struct {
	expected []*struct {
		fun func()
		expectedCalled int
		called int
	}
	then func() *MockunexportedFooWhen
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
}

type MockunexportedFooWhenWithTimes struct {
	*MockunexportedTimes[*MockunexportedFooWhen]
	*MockunexportedFooWhen
}

// Do will execute the provided function and return the result when called
func (_this *MockunexportedFooWhen) Do(do func()) *MockunexportedTimes[*MockunexportedFooWhen] {
	_this.expected[len(_this.expected) -1].fun = do
	return &MockunexportedTimes[*MockunexportedFooWhen] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockunexportedThen: MockunexportedThen[*MockunexportedFooWhen]{ then: _this.then, t: _this.t},
	}
}

type MockunexportedThen [T any] struct {
	then func() T
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
}

// Then continue with another action
func (_this *MockunexportedThen[T]) Then() T {
	_this.t.Helper()
	return _this.then()
}

type MockunexportedTimes[T any] struct {
	expectedCalled *int
	then func() T
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
	MockunexportedThen[T]
}

// Times sets how often the mocked function is expected to be called.
// Test will fail if the number of calls do not match with the expected calls value.
func (_this *MockunexportedTimes[T]) Times(times int) *MockunexportedThen[T] {
	_this.t.Helper()
	*_this.expectedCalled = times
	retVal := &MockunexportedThen[T] { t: _this.t, then: _this.then }
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
func (_this *MockunexportedTimes[T]) AnyTimes() {
	*_this.expectedCalled = -1
}

// Never will fail if the function is ever called.
func (_this *MockunexportedTimes[T]) Never() {
	*_this.expectedCalled = 0
}

