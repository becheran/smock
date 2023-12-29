// DO NOT EDIT
// Code generated by smock 

package testpackage_mock

import (
	fmt "fmt"
	io "io"
	os "os"
	reflect "reflect"
	sync "sync"
	testpackage "github.com/test/testpackage"
)

// NewMockInheritOwn creates a new mock object which implements the corresponding interface.
// All function calls can be mocked with a custom behavior for tests using the WHEN function on the mock object.
// The mock will use the passed in testing.T to report test failures.
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
			for _, c := range v.expected {
				if c.expectedCalled >= 0 && c.expectedCalled != c.called {
					errStr += fmt.Sprintf("\nExpected 'RetType' to be called %d times, but was called %d times.", c.expectedCalled, c.called)
				}
			}
		}
		for _, v := range m.vUseStdType {
			for _, c := range v.expected {
				if c.expectedCalled >= 0 && c.expectedCalled != c.called {
					errStr += fmt.Sprintf("\nExpected 'UseStdType' to be called %d times, but was called %d times.", c.expectedCalled, c.called)
				}
			}
		}
		if errStr != "" {
			t.Helper()
			t.Fatalf(errStr)
		}})
	return m
}

type MockInheritOwn struct {
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
	
	vRetType []*struct{validateArgs func() bool; expected []*struct{fun func() (_r0 testpackage.MyType); expectedCalled int; called int; mutex sync.Mutex}}
	vUseStdType []*struct{validateArgs func(_fi os.FileInfo) bool; expected []*struct{fun func(_fi os.FileInfo) (_r0 io.Reader); expectedCalled int; called int; mutex sync.Mutex}}
}

func (_this *MockInheritOwn) RetType() (_r0 testpackage.MyType) {
	for _, _check := range _this.vRetType {
		if _check.validateArgs == nil || _check.validateArgs() {
			for _ctr, _exp := range _check.expected {
				_exp.mutex.Lock()
				if _exp.expectedCalled <= 0 || _ctr == len(_check.expected) - 1 || _exp.called < _exp.expectedCalled {
					_exp.called++
					_exp.mutex.Unlock()
					return _exp.fun()
				}
				_exp.mutex.Unlock()
			}
		}
	}
	_this.t.Helper()
	_this.unexpectedCall("RetType", )
	return
}

func (_this *MockInheritOwn) UseStdType(_fi os.FileInfo) (_r0 io.Reader) {
	for _, _check := range _this.vUseStdType {
		if _check.validateArgs == nil || _check.validateArgs(_fi) {
			for _ctr, _exp := range _check.expected {
				_exp.mutex.Lock()
				if _exp.expectedCalled <= 0 || _ctr == len(_check.expected) - 1 || _exp.called < _exp.expectedCalled {
					_exp.called++
					_exp.mutex.Unlock()
					return _exp.fun(_fi)
				}
				_exp.mutex.Unlock()
			}
		}
	}
	_this.t.Helper()
	_this.unexpectedCall("UseStdType", _fi)
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
	return &MockInheritOwnWhen {
		m: _this,
	}
}

type MockInheritOwnWhen struct {
	m *MockInheritOwn
}

// Defines the behavior when RetType of the mock is called.
//
// As a default the method is expected to be called once.
// To change this behavior use the Times() method to define how often the function shall be called.
func (_this *MockInheritOwnWhen) RetType() *MockInheritOwnRetTypeWhenWithTimes {
	for _, f := range _this.m.vRetType {
		if f.validateArgs == nil {
			_this.m.t.Helper()
			_this.m.t.Fatalf("Unreachable condition. Call to 'RetType' is already captured by previous WHEN statement.")
		}
	}
	var defaultExpected struct {
		fun func() (_r0 testpackage.MyType)
		expectedCalled int
		called int
		mutex sync.Mutex
	}
	defaultExpected.fun = func() (_r0 testpackage.MyType) { return }
	defaultExpected.expectedCalled = 1
	
	var validator struct {
		validateArgs func() bool
		expected []*struct {
			fun func() (_r0 testpackage.MyType)
			expectedCalled int
			called int
			mutex sync.Mutex
		}
	}
	validator.expected = append(validator.expected, &defaultExpected)
	_this.m.vRetType = append(_this.m.vRetType, &validator)
	var _then func() *MockInheritOwnRetTypeWhen
	_then = func() *MockInheritOwnRetTypeWhen {
		var _newExpected struct {
			fun func() (_r0 testpackage.MyType)
			expectedCalled int
			called int
			mutex sync.Mutex
		}
		_newExpected.fun = func() (_r0 testpackage.MyType) { return }
		_newExpected.expectedCalled = 1
		
		validator.expected = append(validator.expected, &_newExpected)
		return &MockInheritOwnRetTypeWhen {
			expected: validator.expected,
			then: _then,
			t: _this.m.t,
		}
	}
	
	times := &MockInheritOwnTimes[*MockInheritOwnRetTypeWhen] {
		expectedCalled: &validator.expected[0].expectedCalled,
		then: _then,
		t: _this.m.t,
		MockInheritOwnThen: MockInheritOwnThen[*MockInheritOwnRetTypeWhen]{ then: _then, t: _this.m.t},
	}
	return &MockInheritOwnRetTypeWhenWithTimes {
		MockInheritOwnRetTypeWhen: &MockInheritOwnRetTypeWhen {
			expected: validator.expected,
			then: _then,
			t: _this.m.t,
		},
		MockInheritOwnTimes: times,
	}
}

type MockInheritOwnRetTypeWhen struct {
	expected []*struct {
		fun func() (_r0 testpackage.MyType)
		expectedCalled int
		called int
		mutex sync.Mutex
	}
	then func() *MockInheritOwnRetTypeWhen
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
}

type MockInheritOwnRetTypeWhenWithTimes struct {
	*MockInheritOwnTimes[*MockInheritOwnRetTypeWhen]
	*MockInheritOwnRetTypeWhen
}

// Return the provided values when called
func (_this *MockInheritOwnRetTypeWhen) Return(_r0 testpackage.MyType) *MockInheritOwnTimes[*MockInheritOwnRetTypeWhen] {
	_this.expected[len(_this.expected) -1].fun = func() (testpackage.MyType) { return _r0 }
	return &MockInheritOwnTimes[*MockInheritOwnRetTypeWhen] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockInheritOwnThen: MockInheritOwnThen[*MockInheritOwnRetTypeWhen]{ then: _this.then, t: _this.t},
	}
}

// Do will execute the provided function and return the result when called
func (_this *MockInheritOwnRetTypeWhen) Do(do func() (_r0 testpackage.MyType)) *MockInheritOwnTimes[*MockInheritOwnRetTypeWhen] {
	_this.expected[len(_this.expected) -1].fun = do
	return &MockInheritOwnTimes[*MockInheritOwnRetTypeWhen] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockInheritOwnThen: MockInheritOwnThen[*MockInheritOwnRetTypeWhen]{ then: _this.then, t: _this.t},
	}
}

// Defines the behavior when UseStdType of the mock is called.
//
// As a default the method is expected to be called once.
// To change this behavior use the Times() method to define how often the function shall be called.
func (_this *MockInheritOwnWhen) UseStdType() *MockInheritOwnUseStdTypeExpectWithTimes {
	for _, f := range _this.m.vUseStdType {
		if f.validateArgs == nil {
			_this.m.t.Helper()
			_this.m.t.Fatalf("Unreachable condition. Call to 'UseStdType' is already captured by previous WHEN statement.")
		}
	}
	var defaultExpected struct {
		fun func(_fi os.FileInfo) (_r0 io.Reader)
		expectedCalled int
		called int
		mutex sync.Mutex
	}
	defaultExpected.fun = func(_fi os.FileInfo) (_r0 io.Reader) { return }
	defaultExpected.expectedCalled = 1
	
	var validator struct {
		validateArgs func(_fi os.FileInfo) bool
		expected []*struct {
			fun func(_fi os.FileInfo) (_r0 io.Reader)
			expectedCalled int
			called int
			mutex sync.Mutex
		}
	}
	validator.expected = append(validator.expected, &defaultExpected)
	_this.m.vUseStdType = append(_this.m.vUseStdType, &validator)
	var _then func() *MockInheritOwnUseStdTypeWhen
	_then = func() *MockInheritOwnUseStdTypeWhen {
		var _newExpected struct {
			fun func(_fi os.FileInfo) (_r0 io.Reader)
			expectedCalled int
			called int
			mutex sync.Mutex
		}
		_newExpected.fun = func(_fi os.FileInfo) (_r0 io.Reader) { return }
		_newExpected.expectedCalled = 1
		
		validator.expected = append(validator.expected, &_newExpected)
		return &MockInheritOwnUseStdTypeWhen {
			expected: validator.expected,
			then: _then,
			t: _this.m.t,
		}
	}
	
	times := &MockInheritOwnTimes[*MockInheritOwnUseStdTypeWhen] {
		expectedCalled: &validator.expected[0].expectedCalled,
		then: _then,
		t: _this.m.t,
		MockInheritOwnThen: MockInheritOwnThen[*MockInheritOwnUseStdTypeWhen]{ then: _then, t: _this.m.t},
	}
	return &MockInheritOwnUseStdTypeExpectWithTimes {
		MockInheritOwnUseStdTypeExpect: &MockInheritOwnUseStdTypeExpect {
			MockInheritOwnUseStdTypeWhen: &MockInheritOwnUseStdTypeWhen {
				expected: validator.expected,
				then: _then,
				t: _this.m.t,
			},
			validateArgs: &validator.validateArgs,
			times: times,
		},
		MockInheritOwnTimes: times,
	}
}

type MockInheritOwnUseStdTypeExpect struct {
	*MockInheritOwnUseStdTypeWhen
	validateArgs *func(_fi os.FileInfo) bool
	times *MockInheritOwnTimes[*MockInheritOwnUseStdTypeWhen]
}

// Expect will filter for given arguments.
// Each argument is matched with a filter function. Only if all arguments match this mocked function will be called.

// Arguments are either evaluated using the function, or ignored and always true if the function is set to nil.
func (_this *MockInheritOwnUseStdTypeExpect) Expect(_fi func(os.FileInfo) bool) *MockInheritOwnUseStdTypeWhenWithTimes {
	if !(_fi == nil) {
		*_this.validateArgs = func(__fi os.FileInfo) bool {
			return (_fi == nil || _fi(__fi))
		}
	}
	return &MockInheritOwnUseStdTypeWhenWithTimes {
		MockInheritOwnUseStdTypeWhen: _this.MockInheritOwnUseStdTypeWhen,
		MockInheritOwnTimes: _this.times,
	}
}

type MockInheritOwnUseStdTypeExpectWithTimes struct {
	*MockInheritOwnTimes[*MockInheritOwnUseStdTypeWhen]
	*MockInheritOwnUseStdTypeExpect
}

type MockInheritOwnUseStdTypeWhen struct {
	expected []*struct {
		fun func(_fi os.FileInfo) (_r0 io.Reader)
		expectedCalled int
		called int
		mutex sync.Mutex
	}
	then func() *MockInheritOwnUseStdTypeWhen
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
}

type MockInheritOwnUseStdTypeWhenWithTimes struct {
	*MockInheritOwnTimes[*MockInheritOwnUseStdTypeWhen]
	*MockInheritOwnUseStdTypeWhen
}

// Return the provided values when called
func (_this *MockInheritOwnUseStdTypeWhen) Return(_r0 io.Reader) *MockInheritOwnTimes[*MockInheritOwnUseStdTypeWhen] {
	_this.expected[len(_this.expected) -1].fun = func(os.FileInfo) (io.Reader) { return _r0 }
	return &MockInheritOwnTimes[*MockInheritOwnUseStdTypeWhen] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockInheritOwnThen: MockInheritOwnThen[*MockInheritOwnUseStdTypeWhen]{ then: _this.then, t: _this.t},
	}
}

// Do will execute the provided function and return the result when called
func (_this *MockInheritOwnUseStdTypeWhen) Do(do func(_fi os.FileInfo) (_r0 io.Reader)) *MockInheritOwnTimes[*MockInheritOwnUseStdTypeWhen] {
	_this.expected[len(_this.expected) -1].fun = do
	return &MockInheritOwnTimes[*MockInheritOwnUseStdTypeWhen] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockInheritOwnThen: MockInheritOwnThen[*MockInheritOwnUseStdTypeWhen]{ then: _this.then, t: _this.t},
	}
}

type MockInheritOwnThen [T any] struct {
	then func() T
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
}

// Then continue with another action
func (_this *MockInheritOwnThen[T]) Then() T {
	_this.t.Helper()
	return _this.then()
}

type MockInheritOwnTimes[T any] struct {
	expectedCalled *int
	then func() T
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
	MockInheritOwnThen[T]
}

// Times sets how often the mocked function is expected to be called.
// Test will fail if the number of calls do not match with the expected calls value.
func (_this *MockInheritOwnTimes[T]) Times(times int) *MockInheritOwnThen[T] {
	_this.t.Helper()
	*_this.expectedCalled = times
	retVal := &MockInheritOwnThen[T] { t: _this.t, then: _this.then }
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
func (_this *MockInheritOwnTimes[T]) AnyTimes() {
	*_this.expectedCalled = -1
}

// Never will fail if the function is ever called.
func (_this *MockInheritOwnTimes[T]) Never() {
	*_this.expectedCalled = 0
}

