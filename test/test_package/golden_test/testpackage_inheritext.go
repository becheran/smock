// DO NOT EDIT
// Code generated by smock (unknown)

package testpackage_mock

import (
	fmt "fmt"
	reflect "reflect"
	runtime "runtime"
	sync "sync"
)

// NewMockInheritExt creates a new mock object which implements the corresponding interface.
// All function calls can be mocked with a custom behavior for tests using the WHEN function on the mock object.
// The mock will use the passed in testing.T to report test failures.
func NewMockInheritExt(t interface {
	Fatalf(format string, args ...any)
	Helper()
	Cleanup(f func())
}) *MockInheritExt {
	t.Helper()
	m := &MockInheritExt{t: t}
	t.Cleanup(func () {
		errStr := ""
		for _, v := range m.vClose {
			for _, c := range v.expected {
				c.mutex.Lock()
				if c.expectedCalled >= 0 && c.expectedCalled != c.called {
					errStr += fmt.Sprintf("\nExpected 'Close' to be called %d times, but was called %d times. (%s)", c.expectedCalled, c.called, v.location)
				}
				c.mutex.Unlock()
			}
		}
		for _, v := range m.vRead {
			for _, c := range v.expected {
				c.mutex.Lock()
				if c.expectedCalled >= 0 && c.expectedCalled != c.called {
					errStr += fmt.Sprintf("\nExpected 'Read' to be called %d times, but was called %d times. (%s)", c.expectedCalled, c.called, v.location)
				}
				c.mutex.Unlock()
			}
		}
		for _, v := range m.vSeek {
			for _, c := range v.expected {
				c.mutex.Lock()
				if c.expectedCalled >= 0 && c.expectedCalled != c.called {
					errStr += fmt.Sprintf("\nExpected 'Seek' to be called %d times, but was called %d times. (%s)", c.expectedCalled, c.called, v.location)
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

type MockInheritExt struct {
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
	
	vClose []*struct{location string; validateArgs func() bool; expected []*struct{fun func() (_r0 error); expectedCalled int; called int; mutex sync.Mutex}}
	vRead []*struct{location string; validateArgs func(_p []byte) bool; expected []*struct{fun func(_p []byte) (_n int, _err error); expectedCalled int; called int; mutex sync.Mutex}}
	vSeek []*struct{location string; validateArgs func(_offset int64, _whence int) bool; expected []*struct{fun func(_offset int64, _whence int) (_r0 int64, _r1 error); expectedCalled int; called int; mutex sync.Mutex}}
}

func (_this *MockInheritExt) Close() (_r0 error) {
	for _, _check := range _this.vClose {
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
	_this.unexpectedCall("Close", )
	return
}

func (_this *MockInheritExt) Read(_p []byte) (_n int, _err error) {
	for _, _check := range _this.vRead {
		if _check.validateArgs == nil || _check.validateArgs(_p) {
			for _ctr, _exp := range _check.expected {
				_exp.mutex.Lock()
				if _exp.expectedCalled <= 0 || _ctr == len(_check.expected) - 1 || _exp.called < _exp.expectedCalled {
					_exp.called++
					_exp.mutex.Unlock()
					return _exp.fun(_p)
				}
				_exp.mutex.Unlock()
			}
		}
	}
	_this.t.Helper()
	_this.unexpectedCall("Read", _p)
	return
}

func (_this *MockInheritExt) Seek(_offset int64, _whence int) (_r0 int64, _r1 error) {
	for _, _check := range _this.vSeek {
		if _check.validateArgs == nil || _check.validateArgs(_offset, _whence) {
			for _ctr, _exp := range _check.expected {
				_exp.mutex.Lock()
				if _exp.expectedCalled <= 0 || _ctr == len(_check.expected) - 1 || _exp.called < _exp.expectedCalled {
					_exp.called++
					_exp.mutex.Unlock()
					return _exp.fun(_offset, _whence)
				}
				_exp.mutex.Unlock()
			}
		}
	}
	_this.t.Helper()
	_this.unexpectedCall("Seek", _offset, _whence)
	return
}

func (_this *MockInheritExt) unexpectedCall(method string, args ...any) {
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
func (_this *MockInheritExt) WHEN() *MockInheritExtWhen {
	return &MockInheritExtWhen {
		m: _this,
	}
}

type MockInheritExtWhen struct {
	m *MockInheritExt
}

// Defines the behavior when Close of the mock is called.
//
// As a default the method is expected to be called once.
// To change this behavior use the Times() method to define how often the function shall be called.
func (_this *MockInheritExtWhen) Close() *MockInheritExtCloseWhenWithTimes {
	for _, f := range _this.m.vClose {
		if f.validateArgs == nil {
			_this.m.t.Helper()
			_this.m.t.Fatalf("Unreachable condition. Call to 'Close' is already captured by previous WHEN statement.")
		}
	}
	var defaultExpected struct {
		fun func() (_r0 error)
		expectedCalled int
		called int
		mutex sync.Mutex
	}
	defaultExpected.fun = func() (_r0 error) { return }
	defaultExpected.expectedCalled = 1
	
	var validator struct {
		location string
		validateArgs func() bool
		expected []*struct {
			fun func() (_r0 error)
			expectedCalled int
			called int
			mutex sync.Mutex
		}
	}
	if _, file, line, ok := runtime.Caller(1); ok {
		validator.location = fmt.Sprintf("%s:%d", file, line)
	}
	validator.expected = append(validator.expected, &defaultExpected)
	_this.m.vClose = append(_this.m.vClose, &validator)
	var _then func() *MockInheritExtCloseWhen
	_then = func() *MockInheritExtCloseWhen {
		var _newExpected struct {
			fun func() (_r0 error)
			expectedCalled int
			called int
			mutex sync.Mutex
		}
		_newExpected.fun = func() (_r0 error) { return }
		_newExpected.expectedCalled = 1
		
		validator.expected = append(validator.expected, &_newExpected)
		return &MockInheritExtCloseWhen {
			expected: validator.expected,
			then: _then,
			t: _this.m.t,
		}
	}
	
	times := &MockInheritExtTimes[*MockInheritExtCloseWhen] {
		expectedCalled: &validator.expected[0].expectedCalled,
		then: _then,
		t: _this.m.t,
		MockInheritExtThen: MockInheritExtThen[*MockInheritExtCloseWhen]{ then: _then, t: _this.m.t},
	}
	return &MockInheritExtCloseWhenWithTimes {
		MockInheritExtCloseWhen: &MockInheritExtCloseWhen {
			expected: validator.expected,
			then: _then,
			t: _this.m.t,
		},
		MockInheritExtTimes: times,
	}
}

type MockInheritExtCloseWhen struct {
	expected []*struct {
		fun func() (_r0 error)
		expectedCalled int
		called int
		mutex sync.Mutex
	}
	then func() *MockInheritExtCloseWhen
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
}

type MockInheritExtCloseWhenWithTimes struct {
	*MockInheritExtTimes[*MockInheritExtCloseWhen]
	*MockInheritExtCloseWhen
}

// Return the provided values when called
func (_this *MockInheritExtCloseWhen) Return(_r0 error) *MockInheritExtTimes[*MockInheritExtCloseWhen] {
	_this.expected[len(_this.expected) -1].fun = func() (error) { return _r0 }
	return &MockInheritExtTimes[*MockInheritExtCloseWhen] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockInheritExtThen: MockInheritExtThen[*MockInheritExtCloseWhen]{ then: _this.then, t: _this.t},
	}
}

// Do will execute the provided function and return the result when called
func (_this *MockInheritExtCloseWhen) Do(do func() (_r0 error)) *MockInheritExtTimes[*MockInheritExtCloseWhen] {
	_this.expected[len(_this.expected) -1].fun = do
	return &MockInheritExtTimes[*MockInheritExtCloseWhen] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockInheritExtThen: MockInheritExtThen[*MockInheritExtCloseWhen]{ then: _this.then, t: _this.t},
	}
}

// Defines the behavior when Read of the mock is called.
//
// As a default the method is expected to be called once.
// To change this behavior use the Times() method to define how often the function shall be called.
func (_this *MockInheritExtWhen) Read() *MockInheritExtReadExpectWithTimes {
	for _, f := range _this.m.vRead {
		if f.validateArgs == nil {
			_this.m.t.Helper()
			_this.m.t.Fatalf("Unreachable condition. Call to 'Read' is already captured by previous WHEN statement.")
		}
	}
	var defaultExpected struct {
		fun func(_p []byte) (_n int, _err error)
		expectedCalled int
		called int
		mutex sync.Mutex
	}
	defaultExpected.fun = func(_p []byte) (_n int, _err error) { return }
	defaultExpected.expectedCalled = 1
	
	var validator struct {
		location string
		validateArgs func(_p []byte) bool
		expected []*struct {
			fun func(_p []byte) (_n int, _err error)
			expectedCalled int
			called int
			mutex sync.Mutex
		}
	}
	if _, file, line, ok := runtime.Caller(1); ok {
		validator.location = fmt.Sprintf("%s:%d", file, line)
	}
	validator.expected = append(validator.expected, &defaultExpected)
	_this.m.vRead = append(_this.m.vRead, &validator)
	var _then func() *MockInheritExtReadWhen
	_then = func() *MockInheritExtReadWhen {
		var _newExpected struct {
			fun func(_p []byte) (_n int, _err error)
			expectedCalled int
			called int
			mutex sync.Mutex
		}
		_newExpected.fun = func(_p []byte) (_n int, _err error) { return }
		_newExpected.expectedCalled = 1
		
		validator.expected = append(validator.expected, &_newExpected)
		return &MockInheritExtReadWhen {
			expected: validator.expected,
			then: _then,
			t: _this.m.t,
		}
	}
	
	times := &MockInheritExtTimes[*MockInheritExtReadWhen] {
		expectedCalled: &validator.expected[0].expectedCalled,
		then: _then,
		t: _this.m.t,
		MockInheritExtThen: MockInheritExtThen[*MockInheritExtReadWhen]{ then: _then, t: _this.m.t},
	}
	return &MockInheritExtReadExpectWithTimes {
		MockInheritExtReadExpect: &MockInheritExtReadExpect {
			MockInheritExtReadWhen: &MockInheritExtReadWhen {
				expected: validator.expected,
				then: _then,
				t: _this.m.t,
			},
			validateArgs: &validator.validateArgs,
			times: times,
		},
		MockInheritExtTimes: times,
	}
}

type MockInheritExtReadExpect struct {
	*MockInheritExtReadWhen
	validateArgs *func(_p []byte) bool
	times *MockInheritExtTimes[*MockInheritExtReadWhen]
}

// Expect will filter for given arguments.
// Each argument is matched with a filter function. Only if all arguments match this mocked function will be called.

// Arguments are either evaluated using the function, or ignored and always true if the function is set to nil.
func (_this *MockInheritExtReadExpect) Expect(_p func([]byte) bool) *MockInheritExtReadWhenWithTimes {
	if !(_p == nil) {
		*_this.validateArgs = func(__p []byte) bool {
			return (_p == nil || _p(__p))
		}
	}
	return &MockInheritExtReadWhenWithTimes {
		MockInheritExtReadWhen: _this.MockInheritExtReadWhen,
		MockInheritExtTimes: _this.times,
	}
}

type MockInheritExtReadExpectWithTimes struct {
	*MockInheritExtTimes[*MockInheritExtReadWhen]
	*MockInheritExtReadExpect
}

type MockInheritExtReadWhen struct {
	expected []*struct {
		fun func(_p []byte) (_n int, _err error)
		expectedCalled int
		called int
		mutex sync.Mutex
	}
	then func() *MockInheritExtReadWhen
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
}

type MockInheritExtReadWhenWithTimes struct {
	*MockInheritExtTimes[*MockInheritExtReadWhen]
	*MockInheritExtReadWhen
}

// Return the provided values when called
func (_this *MockInheritExtReadWhen) Return(_n int, _err error) *MockInheritExtTimes[*MockInheritExtReadWhen] {
	_this.expected[len(_this.expected) -1].fun = func([]byte) (int, error) { return _n, _err }
	return &MockInheritExtTimes[*MockInheritExtReadWhen] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockInheritExtThen: MockInheritExtThen[*MockInheritExtReadWhen]{ then: _this.then, t: _this.t},
	}
}

// Do will execute the provided function and return the result when called
func (_this *MockInheritExtReadWhen) Do(do func(_p []byte) (_n int, _err error)) *MockInheritExtTimes[*MockInheritExtReadWhen] {
	_this.expected[len(_this.expected) -1].fun = do
	return &MockInheritExtTimes[*MockInheritExtReadWhen] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockInheritExtThen: MockInheritExtThen[*MockInheritExtReadWhen]{ then: _this.then, t: _this.t},
	}
}

// Defines the behavior when Seek of the mock is called.
//
// As a default the method is expected to be called once.
// To change this behavior use the Times() method to define how often the function shall be called.
func (_this *MockInheritExtWhen) Seek() *MockInheritExtSeekExpectWithTimes {
	for _, f := range _this.m.vSeek {
		if f.validateArgs == nil {
			_this.m.t.Helper()
			_this.m.t.Fatalf("Unreachable condition. Call to 'Seek' is already captured by previous WHEN statement.")
		}
	}
	var defaultExpected struct {
		fun func(_offset int64, _whence int) (_r0 int64, _r1 error)
		expectedCalled int
		called int
		mutex sync.Mutex
	}
	defaultExpected.fun = func(_offset int64, _whence int) (_r0 int64, _r1 error) { return }
	defaultExpected.expectedCalled = 1
	
	var validator struct {
		location string
		validateArgs func(_offset int64, _whence int) bool
		expected []*struct {
			fun func(_offset int64, _whence int) (_r0 int64, _r1 error)
			expectedCalled int
			called int
			mutex sync.Mutex
		}
	}
	if _, file, line, ok := runtime.Caller(1); ok {
		validator.location = fmt.Sprintf("%s:%d", file, line)
	}
	validator.expected = append(validator.expected, &defaultExpected)
	_this.m.vSeek = append(_this.m.vSeek, &validator)
	var _then func() *MockInheritExtSeekWhen
	_then = func() *MockInheritExtSeekWhen {
		var _newExpected struct {
			fun func(_offset int64, _whence int) (_r0 int64, _r1 error)
			expectedCalled int
			called int
			mutex sync.Mutex
		}
		_newExpected.fun = func(_offset int64, _whence int) (_r0 int64, _r1 error) { return }
		_newExpected.expectedCalled = 1
		
		validator.expected = append(validator.expected, &_newExpected)
		return &MockInheritExtSeekWhen {
			expected: validator.expected,
			then: _then,
			t: _this.m.t,
		}
	}
	
	times := &MockInheritExtTimes[*MockInheritExtSeekWhen] {
		expectedCalled: &validator.expected[0].expectedCalled,
		then: _then,
		t: _this.m.t,
		MockInheritExtThen: MockInheritExtThen[*MockInheritExtSeekWhen]{ then: _then, t: _this.m.t},
	}
	return &MockInheritExtSeekExpectWithTimes {
		MockInheritExtSeekExpect: &MockInheritExtSeekExpect {
			MockInheritExtSeekWhen: &MockInheritExtSeekWhen {
				expected: validator.expected,
				then: _then,
				t: _this.m.t,
			},
			validateArgs: &validator.validateArgs,
			times: times,
		},
		MockInheritExtTimes: times,
	}
}

type MockInheritExtSeekExpect struct {
	*MockInheritExtSeekWhen
	validateArgs *func(_offset int64, _whence int) bool
	times *MockInheritExtTimes[*MockInheritExtSeekWhen]
}

// Expect will filter for given arguments.
// Each argument is matched with a filter function. Only if all arguments match this mocked function will be called.

// Arguments are either evaluated using the function, or ignored and always true if the function is set to nil.
func (_this *MockInheritExtSeekExpect) Expect(_offset func(int64) bool, _whence func(int) bool) *MockInheritExtSeekWhenWithTimes {
	if !(_offset == nil && _whence == nil) {
		*_this.validateArgs = func(__offset int64, __whence int) bool {
			return (_offset == nil || _offset(__offset)) && (_whence == nil || _whence(__whence))
		}
	}
	return &MockInheritExtSeekWhenWithTimes {
		MockInheritExtSeekWhen: _this.MockInheritExtSeekWhen,
		MockInheritExtTimes: _this.times,
	}
}

type MockInheritExtSeekExpectWithTimes struct {
	*MockInheritExtTimes[*MockInheritExtSeekWhen]
	*MockInheritExtSeekExpect
}

type MockInheritExtSeekWhen struct {
	expected []*struct {
		fun func(_offset int64, _whence int) (_r0 int64, _r1 error)
		expectedCalled int
		called int
		mutex sync.Mutex
	}
	then func() *MockInheritExtSeekWhen
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
}

type MockInheritExtSeekWhenWithTimes struct {
	*MockInheritExtTimes[*MockInheritExtSeekWhen]
	*MockInheritExtSeekWhen
}

// Return the provided values when called
func (_this *MockInheritExtSeekWhen) Return(_r0 int64, _r1 error) *MockInheritExtTimes[*MockInheritExtSeekWhen] {
	_this.expected[len(_this.expected) -1].fun = func(int64, int) (int64, error) { return _r0, _r1 }
	return &MockInheritExtTimes[*MockInheritExtSeekWhen] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockInheritExtThen: MockInheritExtThen[*MockInheritExtSeekWhen]{ then: _this.then, t: _this.t},
	}
}

// Do will execute the provided function and return the result when called
func (_this *MockInheritExtSeekWhen) Do(do func(_offset int64, _whence int) (_r0 int64, _r1 error)) *MockInheritExtTimes[*MockInheritExtSeekWhen] {
	_this.expected[len(_this.expected) -1].fun = do
	return &MockInheritExtTimes[*MockInheritExtSeekWhen] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockInheritExtThen: MockInheritExtThen[*MockInheritExtSeekWhen]{ then: _this.then, t: _this.t},
	}
}

type MockInheritExtThen [T any] struct {
	then func() T
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
}

// Then continue with another action
func (_this *MockInheritExtThen[T]) Then() T {
	_this.t.Helper()
	return _this.then()
}

type MockInheritExtTimes[T any] struct {
	expectedCalled *int
	then func() T
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
	MockInheritExtThen[T]
}

// Times sets how often the mocked function is expected to be called.
// Test will fail if the number of calls do not match with the expected calls value.
func (_this *MockInheritExtTimes[T]) Times(times int) *MockInheritExtThen[T] {
	_this.t.Helper()
	*_this.expectedCalled = times
	retVal := &MockInheritExtThen[T] { t: _this.t, then: _this.then }
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
func (_this *MockInheritExtTimes[T]) AnyTimes() {
	*_this.expectedCalled = -1
}

// Never will fail if the function is ever called.
func (_this *MockInheritExtTimes[T]) Never() {
	*_this.expectedCalled = 0
}

