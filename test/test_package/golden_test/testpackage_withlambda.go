// DO NOT EDIT
// Code generated by smock 

package testpackage_mock

import (
	fmt "fmt"
	reflect "reflect"
	sync "sync"
)

// NewMockWithLambda creates a new mock object which implements the corresponding interface.
// All function calls can be mocked with a custom behavior for tests using the WHEN function on the mock object.   
func NewMockWithLambda[T comparable](t interface {
	Fatalf(format string, args ...any)
	Helper()
	Cleanup(f func())
}) *MockWithLambda[T] {
	t.Helper()
	m := &MockWithLambda[T]{t: t}
	t.Cleanup(func () {
		errStr := ""
		for _, v := range m.vFoo {
			for _, c := range v.expected {
				if c.expectedCalled >= 0 && c.expectedCalled != c.called {
					errStr += fmt.Sprintf("\nExpected 'Foo' to be called %d times, but was called %d times.", c.expectedCalled, c.called)
				}
			}
		}
		for _, v := range m.vBar {
			for _, c := range v.expected {
				if c.expectedCalled >= 0 && c.expectedCalled != c.called {
					errStr += fmt.Sprintf("\nExpected 'Bar' to be called %d times, but was called %d times.", c.expectedCalled, c.called)
				}
			}
		}
		for _, v := range m.vBaz {
			for _, c := range v.expected {
				if c.expectedCalled >= 0 && c.expectedCalled != c.called {
					errStr += fmt.Sprintf("\nExpected 'Baz' to be called %d times, but was called %d times.", c.expectedCalled, c.called)
				}
			}
		}
		if errStr != "" {
			t.Helper()
			t.Fatalf(errStr)
		}})
	return m
}

type MockWithLambda[T comparable] struct {
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
	
	vFoo []*struct{validateArgs func(_a int, _b ...string) bool; expected []*struct{fun func(_a int, _b ...string) (_r0 bool); expectedCalled int; called int; mutex sync.Mutex}}
	vBar []*struct{validateArgs func(_b ...struct{}) bool; expected []*struct{fun func(_b ...struct{}) (_r0 bool); expectedCalled int; called int; mutex sync.Mutex}}
	vBaz []*struct{validateArgs func(_b ...T) bool; expected []*struct{fun func(_b ...T) (_r0 bool); expectedCalled int; called int; mutex sync.Mutex}}
}

func (_this *MockWithLambda[T]) Foo(_a int, _b ...string) (_r0 bool) {
	for _, _check := range _this.vFoo {
		if _check.validateArgs == nil || _check.validateArgs(_a, _b...) {
			for _ctr, _exp := range _check.expected {
				_exp.mutex.Lock()
				if _exp.expectedCalled <= 0 || _ctr == len(_check.expected) - 1 || _exp.called < _exp.expectedCalled {
					_exp.called++
					_exp.mutex.Unlock()
					return _exp.fun(_a, _b...)
				}
				_exp.mutex.Unlock()
			}
		}
	}
	_this.t.Helper()
	_this.unexpectedCall("Foo", _a, _b)
	return
}

func (_this *MockWithLambda[T]) Bar(_b ...struct{}) (_r0 bool) {
	for _, _check := range _this.vBar {
		if _check.validateArgs == nil || _check.validateArgs(_b...) {
			for _ctr, _exp := range _check.expected {
				_exp.mutex.Lock()
				if _exp.expectedCalled <= 0 || _ctr == len(_check.expected) - 1 || _exp.called < _exp.expectedCalled {
					_exp.called++
					_exp.mutex.Unlock()
					return _exp.fun(_b...)
				}
				_exp.mutex.Unlock()
			}
		}
	}
	_this.t.Helper()
	_this.unexpectedCall("Bar", _b)
	return
}

func (_this *MockWithLambda[T]) Baz(_b ...T) (_r0 bool) {
	for _, _check := range _this.vBaz {
		if _check.validateArgs == nil || _check.validateArgs(_b...) {
			for _ctr, _exp := range _check.expected {
				_exp.mutex.Lock()
				if _exp.expectedCalled <= 0 || _ctr == len(_check.expected) - 1 || _exp.called < _exp.expectedCalled {
					_exp.called++
					_exp.mutex.Unlock()
					return _exp.fun(_b...)
				}
				_exp.mutex.Unlock()
			}
		}
	}
	_this.t.Helper()
	_this.unexpectedCall("Baz", _b)
	return
}

func (_this *MockWithLambda[T]) unexpectedCall(method string, args ...any) {
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
func (_this *MockWithLambda[T]) WHEN() *MockWithLambdaWhen[T] {
	return &MockWithLambdaWhen[T] {
		m: _this,
	}
}

type MockWithLambdaWhen[T comparable] struct {
	m *MockWithLambda[T]
}

// Defines the behavior when Foo of the mock is called.
//
// As a default the method is expected to be called once.
// To change this behavior use the Times() method to define how often the function shall be called.
func (_this *MockWithLambdaWhen[T]) Foo() *MockWithLambdaFooExpectWithTimes[T] {
	for _, f := range _this.m.vFoo {
		if f.validateArgs == nil {
			_this.m.t.Helper()
			_this.m.t.Fatalf("Unreachable condition. Call to 'Foo' is already captured by previous WHEN statement.")
		}
	}
	var defaultExpected struct {
		fun func(_a int, _b ...string) (_r0 bool)
		expectedCalled int
		called int
		mutex sync.Mutex
	}
	defaultExpected.fun = func(_a int, _b ...string) (_r0 bool) { return }
	defaultExpected.expectedCalled = 1
	
	var validator struct {
		validateArgs func(_a int, _b ...string) bool
		expected []*struct {
			fun func(_a int, _b ...string) (_r0 bool)
			expectedCalled int
			called int
			mutex sync.Mutex
		}
	}
	validator.expected = append(validator.expected, &defaultExpected)
	_this.m.vFoo = append(_this.m.vFoo, &validator)
	var _then func() *MockWithLambdaFooWhen[T]
	_then = func() *MockWithLambdaFooWhen[T] {
		var _newExpected struct {
			fun func(_a int, _b ...string) (_r0 bool)
			expectedCalled int
			called int
			mutex sync.Mutex
		}
		_newExpected.fun = func(_a int, _b ...string) (_r0 bool) { return }
		_newExpected.expectedCalled = 1
		
		validator.expected = append(validator.expected, &_newExpected)
		return &MockWithLambdaFooWhen[T] {
			expected: validator.expected,
			then: _then,
			t: _this.m.t,
		}
	}
	
	times := &MockWithLambdaTimes[*MockWithLambdaFooWhen[T]] {
		expectedCalled: &validator.expected[0].expectedCalled,
		then: _then,
		t: _this.m.t,
		MockWithLambdaThen: MockWithLambdaThen[*MockWithLambdaFooWhen[T]]{ then: _then, t: _this.m.t},
	}
	return &MockWithLambdaFooExpectWithTimes[T] {
		MockWithLambdaFooExpect: &MockWithLambdaFooExpect[T] {
			MockWithLambdaFooWhen: &MockWithLambdaFooWhen[T] {
				expected: validator.expected,
				then: _then,
				t: _this.m.t,
			},
			validateArgs: &validator.validateArgs,
			times: times,
		},
		MockWithLambdaTimes: times,
	}
}

type MockWithLambdaFooExpect[T comparable] struct {
	*MockWithLambdaFooWhen[T]
	validateArgs *func(_a int, _b ...string) bool
	times *MockWithLambdaTimes[*MockWithLambdaFooWhen[T]]
}

// Expect will filter for given arguments.
// Each argument is matched with a filter function. Only if all arguments match this mocked function will be called.

// Arguments are either evaluated using the function, or ignored and always true if the function is set to nil.
func (_this *MockWithLambdaFooExpect[T]) Expect(_a func(int) bool, _b ...func(string) bool) *MockWithLambdaFooWhenWithTimes[T] {
	if !(_a == nil && len(_b) == 0) {
		*_this.validateArgs = func(__a int, __b ...string) bool {
			for _idx, _val := range __b {
				if _idx >= len(_b) || !(_b[_idx] == nil || _b[_idx](_val)) {
					return false
				}
			}
			return (_a == nil || _a(__a)) && true
		}
	}
	return &MockWithLambdaFooWhenWithTimes[T] {
		MockWithLambdaFooWhen: _this.MockWithLambdaFooWhen,
		MockWithLambdaTimes: _this.times,
	}
}

type MockWithLambdaFooExpectWithTimes[T comparable] struct {
	*MockWithLambdaTimes[*MockWithLambdaFooWhen[T]]
	*MockWithLambdaFooExpect[T]
}

type MockWithLambdaFooWhen[T comparable] struct {
	expected []*struct {
		fun func(_a int, _b ...string) (_r0 bool)
		expectedCalled int
		called int
		mutex sync.Mutex
	}
	then func() *MockWithLambdaFooWhen[T]
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
}

type MockWithLambdaFooWhenWithTimes[T comparable] struct {
	*MockWithLambdaTimes[*MockWithLambdaFooWhen[T]]
	*MockWithLambdaFooWhen[T]
}

// Return the provided values when called
func (_this *MockWithLambdaFooWhen[T]) Return(_r0 bool) *MockWithLambdaTimes[*MockWithLambdaFooWhen[T]] {
	_this.expected[len(_this.expected) -1].fun = func(int, ...string) (bool) { return _r0 }
	return &MockWithLambdaTimes[*MockWithLambdaFooWhen[T]] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockWithLambdaThen: MockWithLambdaThen[*MockWithLambdaFooWhen[T]]{ then: _this.then, t: _this.t},
	}
}

// Do will execute the provided function and return the result when called
func (_this *MockWithLambdaFooWhen[T]) Do(do func(_a int, _b ...string) (_r0 bool)) *MockWithLambdaTimes[*MockWithLambdaFooWhen[T]] {
	_this.expected[len(_this.expected) -1].fun = do
	return &MockWithLambdaTimes[*MockWithLambdaFooWhen[T]] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockWithLambdaThen: MockWithLambdaThen[*MockWithLambdaFooWhen[T]]{ then: _this.then, t: _this.t},
	}
}

// Defines the behavior when Bar of the mock is called.
//
// As a default the method is expected to be called once.
// To change this behavior use the Times() method to define how often the function shall be called.
func (_this *MockWithLambdaWhen[T]) Bar() *MockWithLambdaBarExpectWithTimes[T] {
	for _, f := range _this.m.vBar {
		if f.validateArgs == nil {
			_this.m.t.Helper()
			_this.m.t.Fatalf("Unreachable condition. Call to 'Bar' is already captured by previous WHEN statement.")
		}
	}
	var defaultExpected struct {
		fun func(_b ...struct{}) (_r0 bool)
		expectedCalled int
		called int
		mutex sync.Mutex
	}
	defaultExpected.fun = func(_b ...struct{}) (_r0 bool) { return }
	defaultExpected.expectedCalled = 1
	
	var validator struct {
		validateArgs func(_b ...struct{}) bool
		expected []*struct {
			fun func(_b ...struct{}) (_r0 bool)
			expectedCalled int
			called int
			mutex sync.Mutex
		}
	}
	validator.expected = append(validator.expected, &defaultExpected)
	_this.m.vBar = append(_this.m.vBar, &validator)
	var _then func() *MockWithLambdaBarWhen[T]
	_then = func() *MockWithLambdaBarWhen[T] {
		var _newExpected struct {
			fun func(_b ...struct{}) (_r0 bool)
			expectedCalled int
			called int
			mutex sync.Mutex
		}
		_newExpected.fun = func(_b ...struct{}) (_r0 bool) { return }
		_newExpected.expectedCalled = 1
		
		validator.expected = append(validator.expected, &_newExpected)
		return &MockWithLambdaBarWhen[T] {
			expected: validator.expected,
			then: _then,
			t: _this.m.t,
		}
	}
	
	times := &MockWithLambdaTimes[*MockWithLambdaBarWhen[T]] {
		expectedCalled: &validator.expected[0].expectedCalled,
		then: _then,
		t: _this.m.t,
		MockWithLambdaThen: MockWithLambdaThen[*MockWithLambdaBarWhen[T]]{ then: _then, t: _this.m.t},
	}
	return &MockWithLambdaBarExpectWithTimes[T] {
		MockWithLambdaBarExpect: &MockWithLambdaBarExpect[T] {
			MockWithLambdaBarWhen: &MockWithLambdaBarWhen[T] {
				expected: validator.expected,
				then: _then,
				t: _this.m.t,
			},
			validateArgs: &validator.validateArgs,
			times: times,
		},
		MockWithLambdaTimes: times,
	}
}

type MockWithLambdaBarExpect[T comparable] struct {
	*MockWithLambdaBarWhen[T]
	validateArgs *func(_b ...struct{}) bool
	times *MockWithLambdaTimes[*MockWithLambdaBarWhen[T]]
}

// Expect will filter for given arguments.
// Each argument is matched with a filter function. Only if all arguments match this mocked function will be called.

// Arguments are either evaluated using the function, or ignored and always true if the function is set to nil.
func (_this *MockWithLambdaBarExpect[T]) Expect(_b ...func(struct{}) bool) *MockWithLambdaBarWhenWithTimes[T] {
	if !(len(_b) == 0) {
		*_this.validateArgs = func(__b ...struct{}) bool {
			for _idx, _val := range __b {
				if _idx >= len(_b) || !(_b[_idx] == nil || _b[_idx](_val)) {
					return false
				}
			}
			return true
		}
	}
	return &MockWithLambdaBarWhenWithTimes[T] {
		MockWithLambdaBarWhen: _this.MockWithLambdaBarWhen,
		MockWithLambdaTimes: _this.times,
	}
}

type MockWithLambdaBarExpectWithTimes[T comparable] struct {
	*MockWithLambdaTimes[*MockWithLambdaBarWhen[T]]
	*MockWithLambdaBarExpect[T]
}

type MockWithLambdaBarWhen[T comparable] struct {
	expected []*struct {
		fun func(_b ...struct{}) (_r0 bool)
		expectedCalled int
		called int
		mutex sync.Mutex
	}
	then func() *MockWithLambdaBarWhen[T]
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
}

type MockWithLambdaBarWhenWithTimes[T comparable] struct {
	*MockWithLambdaTimes[*MockWithLambdaBarWhen[T]]
	*MockWithLambdaBarWhen[T]
}

// Return the provided values when called
func (_this *MockWithLambdaBarWhen[T]) Return(_r0 bool) *MockWithLambdaTimes[*MockWithLambdaBarWhen[T]] {
	_this.expected[len(_this.expected) -1].fun = func(...struct{}) (bool) { return _r0 }
	return &MockWithLambdaTimes[*MockWithLambdaBarWhen[T]] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockWithLambdaThen: MockWithLambdaThen[*MockWithLambdaBarWhen[T]]{ then: _this.then, t: _this.t},
	}
}

// Do will execute the provided function and return the result when called
func (_this *MockWithLambdaBarWhen[T]) Do(do func(_b ...struct{}) (_r0 bool)) *MockWithLambdaTimes[*MockWithLambdaBarWhen[T]] {
	_this.expected[len(_this.expected) -1].fun = do
	return &MockWithLambdaTimes[*MockWithLambdaBarWhen[T]] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockWithLambdaThen: MockWithLambdaThen[*MockWithLambdaBarWhen[T]]{ then: _this.then, t: _this.t},
	}
}

// Defines the behavior when Baz of the mock is called.
//
// As a default the method is expected to be called once.
// To change this behavior use the Times() method to define how often the function shall be called.
func (_this *MockWithLambdaWhen[T]) Baz() *MockWithLambdaBazExpectWithTimes[T] {
	for _, f := range _this.m.vBaz {
		if f.validateArgs == nil {
			_this.m.t.Helper()
			_this.m.t.Fatalf("Unreachable condition. Call to 'Baz' is already captured by previous WHEN statement.")
		}
	}
	var defaultExpected struct {
		fun func(_b ...T) (_r0 bool)
		expectedCalled int
		called int
		mutex sync.Mutex
	}
	defaultExpected.fun = func(_b ...T) (_r0 bool) { return }
	defaultExpected.expectedCalled = 1
	
	var validator struct {
		validateArgs func(_b ...T) bool
		expected []*struct {
			fun func(_b ...T) (_r0 bool)
			expectedCalled int
			called int
			mutex sync.Mutex
		}
	}
	validator.expected = append(validator.expected, &defaultExpected)
	_this.m.vBaz = append(_this.m.vBaz, &validator)
	var _then func() *MockWithLambdaBazWhen[T]
	_then = func() *MockWithLambdaBazWhen[T] {
		var _newExpected struct {
			fun func(_b ...T) (_r0 bool)
			expectedCalled int
			called int
			mutex sync.Mutex
		}
		_newExpected.fun = func(_b ...T) (_r0 bool) { return }
		_newExpected.expectedCalled = 1
		
		validator.expected = append(validator.expected, &_newExpected)
		return &MockWithLambdaBazWhen[T] {
			expected: validator.expected,
			then: _then,
			t: _this.m.t,
		}
	}
	
	times := &MockWithLambdaTimes[*MockWithLambdaBazWhen[T]] {
		expectedCalled: &validator.expected[0].expectedCalled,
		then: _then,
		t: _this.m.t,
		MockWithLambdaThen: MockWithLambdaThen[*MockWithLambdaBazWhen[T]]{ then: _then, t: _this.m.t},
	}
	return &MockWithLambdaBazExpectWithTimes[T] {
		MockWithLambdaBazExpect: &MockWithLambdaBazExpect[T] {
			MockWithLambdaBazWhen: &MockWithLambdaBazWhen[T] {
				expected: validator.expected,
				then: _then,
				t: _this.m.t,
			},
			validateArgs: &validator.validateArgs,
			times: times,
		},
		MockWithLambdaTimes: times,
	}
}

type MockWithLambdaBazExpect[T comparable] struct {
	*MockWithLambdaBazWhen[T]
	validateArgs *func(_b ...T) bool
	times *MockWithLambdaTimes[*MockWithLambdaBazWhen[T]]
}

// Expect will filter for given arguments.
// Each argument is matched with a filter function. Only if all arguments match this mocked function will be called.

// Arguments are either evaluated using the function, or ignored and always true if the function is set to nil.
func (_this *MockWithLambdaBazExpect[T]) Expect(_b ...func(T) bool) *MockWithLambdaBazWhenWithTimes[T] {
	if !(len(_b) == 0) {
		*_this.validateArgs = func(__b ...T) bool {
			for _idx, _val := range __b {
				if _idx >= len(_b) || !(_b[_idx] == nil || _b[_idx](_val)) {
					return false
				}
			}
			return true
		}
	}
	return &MockWithLambdaBazWhenWithTimes[T] {
		MockWithLambdaBazWhen: _this.MockWithLambdaBazWhen,
		MockWithLambdaTimes: _this.times,
	}
}

type MockWithLambdaBazExpectWithTimes[T comparable] struct {
	*MockWithLambdaTimes[*MockWithLambdaBazWhen[T]]
	*MockWithLambdaBazExpect[T]
}

type MockWithLambdaBazWhen[T comparable] struct {
	expected []*struct {
		fun func(_b ...T) (_r0 bool)
		expectedCalled int
		called int
		mutex sync.Mutex
	}
	then func() *MockWithLambdaBazWhen[T]
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
}

type MockWithLambdaBazWhenWithTimes[T comparable] struct {
	*MockWithLambdaTimes[*MockWithLambdaBazWhen[T]]
	*MockWithLambdaBazWhen[T]
}

// Return the provided values when called
func (_this *MockWithLambdaBazWhen[T]) Return(_r0 bool) *MockWithLambdaTimes[*MockWithLambdaBazWhen[T]] {
	_this.expected[len(_this.expected) -1].fun = func(...T) (bool) { return _r0 }
	return &MockWithLambdaTimes[*MockWithLambdaBazWhen[T]] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockWithLambdaThen: MockWithLambdaThen[*MockWithLambdaBazWhen[T]]{ then: _this.then, t: _this.t},
	}
}

// Do will execute the provided function and return the result when called
func (_this *MockWithLambdaBazWhen[T]) Do(do func(_b ...T) (_r0 bool)) *MockWithLambdaTimes[*MockWithLambdaBazWhen[T]] {
	_this.expected[len(_this.expected) -1].fun = do
	return &MockWithLambdaTimes[*MockWithLambdaBazWhen[T]] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockWithLambdaThen: MockWithLambdaThen[*MockWithLambdaBazWhen[T]]{ then: _this.then, t: _this.t},
	}
}

type MockWithLambdaThen [T any] struct {
	then func() T
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
}

// Then continue with another action
func (_this *MockWithLambdaThen[T]) Then() T {
	_this.t.Helper()
	return _this.then()
}

type MockWithLambdaTimes[T any] struct {
	expectedCalled *int
	then func() T
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
	MockWithLambdaThen[T]
}

// Times sets how often the mocked function is expected to be called.
// Test will fail if the number of calls do not match with the expected calls value.
func (_this *MockWithLambdaTimes[T]) Times(times int) *MockWithLambdaThen[T] {
	_this.t.Helper()
	*_this.expectedCalled = times
	retVal := &MockWithLambdaThen[T] { t: _this.t, then: _this.then }
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
func (_this *MockWithLambdaTimes[T]) AnyTimes() {
	*_this.expectedCalled = -1
}

// Never will fail if the function is ever called.
func (_this *MockWithLambdaTimes[T]) Never() {
	*_this.expectedCalled = 0
}

