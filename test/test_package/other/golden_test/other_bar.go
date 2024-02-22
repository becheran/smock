// DO NOT EDIT
// Code generated by smock (unknown)

package other_mock

import (
	fmt "fmt"
	other "github.com/test/testpackage/other"
	reflect "reflect"
	runtime "runtime"
	sync "sync"
)

// NewMockBar creates a new mock object which implements the corresponding interface.
// All function calls can be mocked with a custom behavior for tests using the WHEN function on the mock object.
// The mock will use the passed in testing.T to report test failures.
func NewMockBar(t interface {
	Fatalf(format string, args ...any)
	Helper()
	Cleanup(f func())
}) *MockBar {
	t.Helper()
	m := &MockBar{t: t}
	t.Cleanup(func () {
		errStr := ""
		for _, v := range m.vDo {
			for _, c := range v.expected {
				c.mutex.Lock()
				if c.expectedCalled >= 0 && c.expectedCalled != c.called {
					errStr += fmt.Sprintf("\nExpected 'Do' to be called %d times, but was called %d times. (%s)", c.expectedCalled, c.called, v.location)
				}
				c.mutex.Unlock()
			}
		}
		if errStr != "" {
			t.Helper()
			t.Fatalf(errStr)
		}})
	return m
}

type MockBar struct {
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
	
	vDo []*struct{location string; validateArgs func(_i0 func(other.Custom) other.Custom) bool; expected []*struct{fun func(_i0 func(other.Custom) other.Custom) (_r0 other.Custom); expectedCalled int; called int; mutex sync.Mutex}}
}

func (_this *MockBar) Do(_i0 func(other.Custom) other.Custom) (_r0 other.Custom) {
	for _, _check := range _this.vDo {
		if _check.validateArgs == nil || _check.validateArgs(_i0) {
			for _ctr, _exp := range _check.expected {
				_exp.mutex.Lock()
				if _exp.expectedCalled <= 0 || _ctr == len(_check.expected) - 1 || _exp.called < _exp.expectedCalled {
					_exp.called++
					_exp.mutex.Unlock()
					return _exp.fun(_i0)
				}
				_exp.mutex.Unlock()
			}
		}
	}
	_this.t.Helper()
	_this.unexpectedCall("Do", _i0)
	return
}

func (_this *MockBar) unexpectedCall(method string, args ...any) {
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
func (_this *MockBar) WHEN() *MockBarWhen {
	return &MockBarWhen {
		m: _this,
	}
}

type MockBarWhen struct {
	m *MockBar
}

// Defines the behavior when Do of the mock is called.
//
// As a default the method is expected to be called once.
// To change this behavior use the Times() method to define how often the function shall be called.
func (_this *MockBarWhen) Do() *MockBarDoExpectWithTimes {
	for _, f := range _this.m.vDo {
		if f.validateArgs == nil {
			_this.m.t.Helper()
			_this.m.t.Fatalf("Unreachable condition. Call to 'Do' is already captured by previous WHEN statement.")
		}
	}
	var defaultExpected struct {
		fun func(_i0 func(other.Custom) other.Custom) (_r0 other.Custom)
		expectedCalled int
		called int
		mutex sync.Mutex
	}
	defaultExpected.fun = func(_i0 func(other.Custom) other.Custom) (_r0 other.Custom) { return }
	defaultExpected.expectedCalled = 1
	
	var validator struct {
		location string
		validateArgs func(_i0 func(other.Custom) other.Custom) bool
		expected []*struct {
			fun func(_i0 func(other.Custom) other.Custom) (_r0 other.Custom)
			expectedCalled int
			called int
			mutex sync.Mutex
		}
	}
	if _, file, line, ok := runtime.Caller(1); ok {
		validator.location = fmt.Sprintf("%s:%d", file, line)
	}
	validator.expected = append(validator.expected, &defaultExpected)
	_this.m.vDo = append(_this.m.vDo, &validator)
	var _then func() *MockBarDoWhen
	_then = func() *MockBarDoWhen {
		var _newExpected struct {
			fun func(_i0 func(other.Custom) other.Custom) (_r0 other.Custom)
			expectedCalled int
			called int
			mutex sync.Mutex
		}
		_newExpected.fun = func(_i0 func(other.Custom) other.Custom) (_r0 other.Custom) { return }
		_newExpected.expectedCalled = 1
		
		validator.expected = append(validator.expected, &_newExpected)
		return &MockBarDoWhen {
			expected: validator.expected,
			then: _then,
			t: _this.m.t,
		}
	}
	
	times := &MockBarTimes[*MockBarDoWhen] {
		expectedCalled: &validator.expected[0].expectedCalled,
		then: _then,
		t: _this.m.t,
		MockBarThen: MockBarThen[*MockBarDoWhen]{ then: _then, t: _this.m.t},
	}
	return &MockBarDoExpectWithTimes {
		MockBarDoExpect: &MockBarDoExpect {
			MockBarDoWhen: &MockBarDoWhen {
				expected: validator.expected,
				then: _then,
				t: _this.m.t,
			},
			validateArgs: &validator.validateArgs,
			times: times,
		},
		MockBarTimes: times,
	}
}

type MockBarDoExpect struct {
	*MockBarDoWhen
	validateArgs *func(_i0 func(other.Custom) other.Custom) bool
	times *MockBarTimes[*MockBarDoWhen]
}

// Expect will filter for given arguments.
// Each argument is matched with a filter function. Only if all arguments match this mocked function will be called.

// Arguments are either evaluated using the function, or ignored and always true if the function is set to nil.
func (_this *MockBarDoExpect) Expect(_0 func(func(other.Custom) other.Custom) bool) *MockBarDoWhenWithTimes {
	if !(_0 == nil) {
		*_this.validateArgs = func(__i0 func(other.Custom) other.Custom) bool {
			return (_0 == nil || _0(__i0))
		}
	}
	return &MockBarDoWhenWithTimes {
		MockBarDoWhen: _this.MockBarDoWhen,
		MockBarTimes: _this.times,
	}
}

type MockBarDoExpectWithTimes struct {
	*MockBarTimes[*MockBarDoWhen]
	*MockBarDoExpect
}

type MockBarDoWhen struct {
	expected []*struct {
		fun func(_i0 func(other.Custom) other.Custom) (_r0 other.Custom)
		expectedCalled int
		called int
		mutex sync.Mutex
	}
	then func() *MockBarDoWhen
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
}

type MockBarDoWhenWithTimes struct {
	*MockBarTimes[*MockBarDoWhen]
	*MockBarDoWhen
}

// Return the provided values when called
func (_this *MockBarDoWhen) Return(_r0 other.Custom) *MockBarTimes[*MockBarDoWhen] {
	_this.expected[len(_this.expected) -1].fun = func(func(other.Custom) other.Custom) (other.Custom) { return _r0 }
	return &MockBarTimes[*MockBarDoWhen] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockBarThen: MockBarThen[*MockBarDoWhen]{ then: _this.then, t: _this.t},
	}
}

// Do will execute the provided function and return the result when called
func (_this *MockBarDoWhen) Do(do func(_i0 func(other.Custom) other.Custom) (_r0 other.Custom)) *MockBarTimes[*MockBarDoWhen] {
	_this.expected[len(_this.expected) -1].fun = do
	return &MockBarTimes[*MockBarDoWhen] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockBarThen: MockBarThen[*MockBarDoWhen]{ then: _this.then, t: _this.t},
	}
}

type MockBarThen [T any] struct {
	then func() T
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
}

// Then continue with another action
func (_this *MockBarThen[T]) Then() T {
	_this.t.Helper()
	return _this.then()
}

type MockBarTimes[T any] struct {
	expectedCalled *int
	then func() T
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
	MockBarThen[T]
}

// Times sets how often the mocked function is expected to be called.
// Test will fail if the number of calls do not match with the expected calls value.
func (_this *MockBarTimes[T]) Times(times int) *MockBarThen[T] {
	_this.t.Helper()
	*_this.expectedCalled = times
	retVal := &MockBarThen[T] { t: _this.t, then: _this.then }
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
func (_this *MockBarTimes[T]) AnyTimes() {
	*_this.expectedCalled = -1
}

// Never will fail if the function is ever called.
func (_this *MockBarTimes[T]) Never() {
	*_this.expectedCalled = 0
}

