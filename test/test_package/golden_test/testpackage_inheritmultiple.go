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

// MockInheritMultiple must implement interface testpackage.InheritMultiple
var _ testpackage.InheritMultiple = &MockInheritMultiple{}

// NewMockInheritMultiple creates a new mock object which implements the corresponding interface.
// All function calls can be mocked with a custom behavior for tests using the WHEN function on the mock object.   
func NewMockInheritMultiple(t interface {
	Fatalf(format string, args ...any)
	Helper()
	Cleanup(f func())
}) *MockInheritMultiple {
	t.Helper()
	m := &MockInheritMultiple{t: t}
	t.Cleanup(func () {
		errStr := ""
		for _, v := range m.vOwn {
			for _, c := range v.expected {
				if c.expectedCalled >= 0 && c.expectedCalled != c.called {
					errStr += fmt.Sprintf("\nExpected 'Own' to be called %d times, but was called %d times.", c.expectedCalled, c.called)
				}
			}
		}
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
		for _, v := range m.vClose {
			for _, c := range v.expected {
				if c.expectedCalled >= 0 && c.expectedCalled != c.called {
					errStr += fmt.Sprintf("\nExpected 'Close' to be called %d times, but was called %d times.", c.expectedCalled, c.called)
				}
			}
		}
		for _, v := range m.vRead {
			for _, c := range v.expected {
				if c.expectedCalled >= 0 && c.expectedCalled != c.called {
					errStr += fmt.Sprintf("\nExpected 'Read' to be called %d times, but was called %d times.", c.expectedCalled, c.called)
				}
			}
		}
		for _, v := range m.vSeek {
			for _, c := range v.expected {
				if c.expectedCalled >= 0 && c.expectedCalled != c.called {
					errStr += fmt.Sprintf("\nExpected 'Seek' to be called %d times, but was called %d times.", c.expectedCalled, c.called)
				}
			}
		}
		if errStr != "" {
			t.Helper()
			t.Fatalf(errStr)
		}})
	return m
}

type MockInheritMultiple struct {
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
	
	vOwn []*struct{validateArgs func(i0 int, i1 string) bool; expected []*struct{fun func(i0 int, i1 string) (r0 int, r1 string); expectedCalled int; called int; mutex sync.Mutex}}
	vRetType []*struct{validateArgs func() bool; expected []*struct{fun func() (r0 testpackage.MyType); expectedCalled int; called int; mutex sync.Mutex}}
	vUseStdType []*struct{validateArgs func(fi os.FileInfo) bool; expected []*struct{fun func(fi os.FileInfo) (r0 io.Reader); expectedCalled int; called int; mutex sync.Mutex}}
	vClose []*struct{validateArgs func() bool; expected []*struct{fun func() (r0 error); expectedCalled int; called int; mutex sync.Mutex}}
	vRead []*struct{validateArgs func(p []byte) bool; expected []*struct{fun func(p []byte) (n int, err error); expectedCalled int; called int; mutex sync.Mutex}}
	vSeek []*struct{validateArgs func(offset int64, whence int) bool; expected []*struct{fun func(offset int64, whence int) (r0 int64, r1 error); expectedCalled int; called int; mutex sync.Mutex}}
}

func (_this *MockInheritMultiple) Own(i0 int, i1 string) (r0 int, r1 string) {
	for _, _check := range _this.vOwn {
		if _check.validateArgs == nil || _check.validateArgs(i0, i1) {
			for _ctr, _exp := range _check.expected {
				_exp.mutex.Lock()
				if _exp.expectedCalled <= 0 || _ctr == len(_check.expected) - 1 || _exp.called < _exp.expectedCalled {
					_exp.called++
					_exp.mutex.Unlock()
					return _exp.fun(i0, i1)
				}
				_exp.mutex.Unlock()
			}
		}
	}
	_this.t.Helper()
	_this.unexpectedCall("Own", i0, i1)
	return
}

func (_this *MockInheritMultiple) RetType() (r0 testpackage.MyType) {
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

func (_this *MockInheritMultiple) UseStdType(fi os.FileInfo) (r0 io.Reader) {
	for _, _check := range _this.vUseStdType {
		if _check.validateArgs == nil || _check.validateArgs(fi) {
			for _ctr, _exp := range _check.expected {
				_exp.mutex.Lock()
				if _exp.expectedCalled <= 0 || _ctr == len(_check.expected) - 1 || _exp.called < _exp.expectedCalled {
					_exp.called++
					_exp.mutex.Unlock()
					return _exp.fun(fi)
				}
				_exp.mutex.Unlock()
			}
		}
	}
	_this.t.Helper()
	_this.unexpectedCall("UseStdType", fi)
	return
}

func (_this *MockInheritMultiple) Close() (r0 error) {
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

func (_this *MockInheritMultiple) Read(p []byte) (n int, err error) {
	for _, _check := range _this.vRead {
		if _check.validateArgs == nil || _check.validateArgs(p) {
			for _ctr, _exp := range _check.expected {
				_exp.mutex.Lock()
				if _exp.expectedCalled <= 0 || _ctr == len(_check.expected) - 1 || _exp.called < _exp.expectedCalled {
					_exp.called++
					_exp.mutex.Unlock()
					return _exp.fun(p)
				}
				_exp.mutex.Unlock()
			}
		}
	}
	_this.t.Helper()
	_this.unexpectedCall("Read", p)
	return
}

func (_this *MockInheritMultiple) Seek(offset int64, whence int) (r0 int64, r1 error) {
	for _, _check := range _this.vSeek {
		if _check.validateArgs == nil || _check.validateArgs(offset, whence) {
			for _ctr, _exp := range _check.expected {
				_exp.mutex.Lock()
				if _exp.expectedCalled <= 0 || _ctr == len(_check.expected) - 1 || _exp.called < _exp.expectedCalled {
					_exp.called++
					_exp.mutex.Unlock()
					return _exp.fun(offset, whence)
				}
				_exp.mutex.Unlock()
			}
		}
	}
	_this.t.Helper()
	_this.unexpectedCall("Seek", offset, whence)
	return
}

func (_this *MockInheritMultiple) unexpectedCall(method string, args ...any) {
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
func (_this *MockInheritMultiple) WHEN() *MockInheritMultipleWhen {
	return &MockInheritMultipleWhen {
		m: _this,
	}
}

type MockInheritMultipleWhen struct {
	m *MockInheritMultiple
}

// Defines the behavior when Own of the mock is called.
//
// As a default the method is expected to be called once.
// To change this behavior use the Times() method to define how often the function shall be called.
func (_this *MockInheritMultipleWhen) Own() *MockInheritMultipleOwnExpectWithTimes {
	for _, f := range _this.m.vOwn {
		if f.validateArgs == nil {
			_this.m.t.Helper()
			_this.m.t.Fatalf("Unreachable condition. Call to 'Own' is already captured by previous WHEN statement.")
		}
	}
	var defaultExpected struct {
		fun func(i0 int, i1 string) (r0 int, r1 string)
		expectedCalled int
		called int
		mutex sync.Mutex
	}
	defaultExpected.fun = func(i0 int, i1 string) (r0 int, r1 string) { return }
	defaultExpected.expectedCalled = 1
	
	var validator struct {
		validateArgs func(i0 int, i1 string) bool
		expected []*struct {
			fun func(i0 int, i1 string) (r0 int, r1 string)
			expectedCalled int
			called int
			mutex sync.Mutex
		}
	}
	validator.expected = append(validator.expected, &defaultExpected)
	_this.m.vOwn = append(_this.m.vOwn, &validator)
	var _then func() *MockInheritMultipleOwnWhen
	_then = func() *MockInheritMultipleOwnWhen {
		var _newExpected struct {
			fun func(i0 int, i1 string) (r0 int, r1 string)
			expectedCalled int
			called int
			mutex sync.Mutex
		}
		_newExpected.fun = func(i0 int, i1 string) (r0 int, r1 string) { return }
		_newExpected.expectedCalled = 1
		
		validator.expected = append(validator.expected, &_newExpected)
		return &MockInheritMultipleOwnWhen {
			expected: validator.expected,
			then: _then,
			t: _this.m.t,
		}
	}
	
	times := &MockInheritMultipleTimes[*MockInheritMultipleOwnWhen] {
		expectedCalled: &validator.expected[0].expectedCalled,
		then: _then,
		t: _this.m.t,
		MockInheritMultipleThen: MockInheritMultipleThen[*MockInheritMultipleOwnWhen]{ then: _then, t: _this.m.t},
	}
	return &MockInheritMultipleOwnExpectWithTimes {
		MockInheritMultipleOwnExpect: &MockInheritMultipleOwnExpect {
			MockInheritMultipleOwnWhen: &MockInheritMultipleOwnWhen {
				expected: validator.expected,
				then: _then,
				t: _this.m.t,
			},
			validateArgs: &validator.validateArgs,
			times: times,
		},
		MockInheritMultipleTimes: times,
	}
}

type MockInheritMultipleOwnExpect struct {
	*MockInheritMultipleOwnWhen
	validateArgs *func(i0 int, i1 string) bool
	times *MockInheritMultipleTimes[*MockInheritMultipleOwnWhen]
}

// Expect will filter for given arguments.
// Each argument is matched with a filter function. Only if all arguments match this mocked function will be called.

// Arguments are either evaluated using the function, or ignored and always true if the function is set to nil.
func (_this *MockInheritMultipleOwnExpect) Expect(_0 func(int) bool, _1 func(string) bool) *MockInheritMultipleOwnWhenWithTimes {
	if !(_0 == nil && _1 == nil) {
		*_this.validateArgs = func(_i0 int, _i1 string) bool {
			return (_0 == nil || _0(_i0)) && (_1 == nil || _1(_i1))
		}
	}
	return &MockInheritMultipleOwnWhenWithTimes {
		MockInheritMultipleOwnWhen: _this.MockInheritMultipleOwnWhen,
		MockInheritMultipleTimes: _this.times,
	}
}

type MockInheritMultipleOwnExpectWithTimes struct {
	*MockInheritMultipleTimes[*MockInheritMultipleOwnWhen]
	*MockInheritMultipleOwnExpect
}

type MockInheritMultipleOwnWhen struct {
	expected []*struct {
		fun func(i0 int, i1 string) (r0 int, r1 string)
		expectedCalled int
		called int
		mutex sync.Mutex
	}
	then func() *MockInheritMultipleOwnWhen
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
}

type MockInheritMultipleOwnWhenWithTimes struct {
	*MockInheritMultipleTimes[*MockInheritMultipleOwnWhen]
	*MockInheritMultipleOwnWhen
}

// Return the provided values when called
func (_this *MockInheritMultipleOwnWhen) Return(r0 int, r1 string) *MockInheritMultipleTimes[*MockInheritMultipleOwnWhen] {
	_this.expected[len(_this.expected) -1].fun = func(int, string) (int, string) { return r0, r1 }
	return &MockInheritMultipleTimes[*MockInheritMultipleOwnWhen] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockInheritMultipleThen: MockInheritMultipleThen[*MockInheritMultipleOwnWhen]{ then: _this.then, t: _this.t},
	}
}

// Do will execute the provided function and return the result when called
func (_this *MockInheritMultipleOwnWhen) Do(do func(i0 int, i1 string) (r0 int, r1 string)) *MockInheritMultipleTimes[*MockInheritMultipleOwnWhen] {
	_this.expected[len(_this.expected) -1].fun = do
	return &MockInheritMultipleTimes[*MockInheritMultipleOwnWhen] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockInheritMultipleThen: MockInheritMultipleThen[*MockInheritMultipleOwnWhen]{ then: _this.then, t: _this.t},
	}
}

// Defines the behavior when RetType of the mock is called.
//
// As a default the method is expected to be called once.
// To change this behavior use the Times() method to define how often the function shall be called.
func (_this *MockInheritMultipleWhen) RetType() *MockInheritMultipleRetTypeWhenWithTimes {
	for _, f := range _this.m.vRetType {
		if f.validateArgs == nil {
			_this.m.t.Helper()
			_this.m.t.Fatalf("Unreachable condition. Call to 'RetType' is already captured by previous WHEN statement.")
		}
	}
	var defaultExpected struct {
		fun func() (r0 testpackage.MyType)
		expectedCalled int
		called int
		mutex sync.Mutex
	}
	defaultExpected.fun = func() (r0 testpackage.MyType) { return }
	defaultExpected.expectedCalled = 1
	
	var validator struct {
		validateArgs func() bool
		expected []*struct {
			fun func() (r0 testpackage.MyType)
			expectedCalled int
			called int
			mutex sync.Mutex
		}
	}
	validator.expected = append(validator.expected, &defaultExpected)
	_this.m.vRetType = append(_this.m.vRetType, &validator)
	var _then func() *MockInheritMultipleRetTypeWhen
	_then = func() *MockInheritMultipleRetTypeWhen {
		var _newExpected struct {
			fun func() (r0 testpackage.MyType)
			expectedCalled int
			called int
			mutex sync.Mutex
		}
		_newExpected.fun = func() (r0 testpackage.MyType) { return }
		_newExpected.expectedCalled = 1
		
		validator.expected = append(validator.expected, &_newExpected)
		return &MockInheritMultipleRetTypeWhen {
			expected: validator.expected,
			then: _then,
			t: _this.m.t,
		}
	}
	
	times := &MockInheritMultipleTimes[*MockInheritMultipleRetTypeWhen] {
		expectedCalled: &validator.expected[0].expectedCalled,
		then: _then,
		t: _this.m.t,
		MockInheritMultipleThen: MockInheritMultipleThen[*MockInheritMultipleRetTypeWhen]{ then: _then, t: _this.m.t},
	}
	return &MockInheritMultipleRetTypeWhenWithTimes {
		MockInheritMultipleRetTypeWhen: &MockInheritMultipleRetTypeWhen {
			expected: validator.expected,
			then: _then,
			t: _this.m.t,
		},
		MockInheritMultipleTimes: times,
	}
}

type MockInheritMultipleRetTypeWhen struct {
	expected []*struct {
		fun func() (r0 testpackage.MyType)
		expectedCalled int
		called int
		mutex sync.Mutex
	}
	then func() *MockInheritMultipleRetTypeWhen
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
}

type MockInheritMultipleRetTypeWhenWithTimes struct {
	*MockInheritMultipleTimes[*MockInheritMultipleRetTypeWhen]
	*MockInheritMultipleRetTypeWhen
}

// Return the provided values when called
func (_this *MockInheritMultipleRetTypeWhen) Return(r0 testpackage.MyType) *MockInheritMultipleTimes[*MockInheritMultipleRetTypeWhen] {
	_this.expected[len(_this.expected) -1].fun = func() (testpackage.MyType) { return r0 }
	return &MockInheritMultipleTimes[*MockInheritMultipleRetTypeWhen] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockInheritMultipleThen: MockInheritMultipleThen[*MockInheritMultipleRetTypeWhen]{ then: _this.then, t: _this.t},
	}
}

// Do will execute the provided function and return the result when called
func (_this *MockInheritMultipleRetTypeWhen) Do(do func() (r0 testpackage.MyType)) *MockInheritMultipleTimes[*MockInheritMultipleRetTypeWhen] {
	_this.expected[len(_this.expected) -1].fun = do
	return &MockInheritMultipleTimes[*MockInheritMultipleRetTypeWhen] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockInheritMultipleThen: MockInheritMultipleThen[*MockInheritMultipleRetTypeWhen]{ then: _this.then, t: _this.t},
	}
}

// Defines the behavior when UseStdType of the mock is called.
//
// As a default the method is expected to be called once.
// To change this behavior use the Times() method to define how often the function shall be called.
func (_this *MockInheritMultipleWhen) UseStdType() *MockInheritMultipleUseStdTypeExpectWithTimes {
	for _, f := range _this.m.vUseStdType {
		if f.validateArgs == nil {
			_this.m.t.Helper()
			_this.m.t.Fatalf("Unreachable condition. Call to 'UseStdType' is already captured by previous WHEN statement.")
		}
	}
	var defaultExpected struct {
		fun func(fi os.FileInfo) (r0 io.Reader)
		expectedCalled int
		called int
		mutex sync.Mutex
	}
	defaultExpected.fun = func(fi os.FileInfo) (r0 io.Reader) { return }
	defaultExpected.expectedCalled = 1
	
	var validator struct {
		validateArgs func(fi os.FileInfo) bool
		expected []*struct {
			fun func(fi os.FileInfo) (r0 io.Reader)
			expectedCalled int
			called int
			mutex sync.Mutex
		}
	}
	validator.expected = append(validator.expected, &defaultExpected)
	_this.m.vUseStdType = append(_this.m.vUseStdType, &validator)
	var _then func() *MockInheritMultipleUseStdTypeWhen
	_then = func() *MockInheritMultipleUseStdTypeWhen {
		var _newExpected struct {
			fun func(fi os.FileInfo) (r0 io.Reader)
			expectedCalled int
			called int
			mutex sync.Mutex
		}
		_newExpected.fun = func(fi os.FileInfo) (r0 io.Reader) { return }
		_newExpected.expectedCalled = 1
		
		validator.expected = append(validator.expected, &_newExpected)
		return &MockInheritMultipleUseStdTypeWhen {
			expected: validator.expected,
			then: _then,
			t: _this.m.t,
		}
	}
	
	times := &MockInheritMultipleTimes[*MockInheritMultipleUseStdTypeWhen] {
		expectedCalled: &validator.expected[0].expectedCalled,
		then: _then,
		t: _this.m.t,
		MockInheritMultipleThen: MockInheritMultipleThen[*MockInheritMultipleUseStdTypeWhen]{ then: _then, t: _this.m.t},
	}
	return &MockInheritMultipleUseStdTypeExpectWithTimes {
		MockInheritMultipleUseStdTypeExpect: &MockInheritMultipleUseStdTypeExpect {
			MockInheritMultipleUseStdTypeWhen: &MockInheritMultipleUseStdTypeWhen {
				expected: validator.expected,
				then: _then,
				t: _this.m.t,
			},
			validateArgs: &validator.validateArgs,
			times: times,
		},
		MockInheritMultipleTimes: times,
	}
}

type MockInheritMultipleUseStdTypeExpect struct {
	*MockInheritMultipleUseStdTypeWhen
	validateArgs *func(fi os.FileInfo) bool
	times *MockInheritMultipleTimes[*MockInheritMultipleUseStdTypeWhen]
}

// Expect will filter for given arguments.
// Each argument is matched with a filter function. Only if all arguments match this mocked function will be called.

// Arguments are either evaluated using the function, or ignored and always true if the function is set to nil.
func (_this *MockInheritMultipleUseStdTypeExpect) Expect(fi func(os.FileInfo) bool) *MockInheritMultipleUseStdTypeWhenWithTimes {
	if !(fi == nil) {
		*_this.validateArgs = func(_fi os.FileInfo) bool {
			return (fi == nil || fi(_fi))
		}
	}
	return &MockInheritMultipleUseStdTypeWhenWithTimes {
		MockInheritMultipleUseStdTypeWhen: _this.MockInheritMultipleUseStdTypeWhen,
		MockInheritMultipleTimes: _this.times,
	}
}

type MockInheritMultipleUseStdTypeExpectWithTimes struct {
	*MockInheritMultipleTimes[*MockInheritMultipleUseStdTypeWhen]
	*MockInheritMultipleUseStdTypeExpect
}

type MockInheritMultipleUseStdTypeWhen struct {
	expected []*struct {
		fun func(fi os.FileInfo) (r0 io.Reader)
		expectedCalled int
		called int
		mutex sync.Mutex
	}
	then func() *MockInheritMultipleUseStdTypeWhen
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
}

type MockInheritMultipleUseStdTypeWhenWithTimes struct {
	*MockInheritMultipleTimes[*MockInheritMultipleUseStdTypeWhen]
	*MockInheritMultipleUseStdTypeWhen
}

// Return the provided values when called
func (_this *MockInheritMultipleUseStdTypeWhen) Return(r0 io.Reader) *MockInheritMultipleTimes[*MockInheritMultipleUseStdTypeWhen] {
	_this.expected[len(_this.expected) -1].fun = func(os.FileInfo) (io.Reader) { return r0 }
	return &MockInheritMultipleTimes[*MockInheritMultipleUseStdTypeWhen] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockInheritMultipleThen: MockInheritMultipleThen[*MockInheritMultipleUseStdTypeWhen]{ then: _this.then, t: _this.t},
	}
}

// Do will execute the provided function and return the result when called
func (_this *MockInheritMultipleUseStdTypeWhen) Do(do func(fi os.FileInfo) (r0 io.Reader)) *MockInheritMultipleTimes[*MockInheritMultipleUseStdTypeWhen] {
	_this.expected[len(_this.expected) -1].fun = do
	return &MockInheritMultipleTimes[*MockInheritMultipleUseStdTypeWhen] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockInheritMultipleThen: MockInheritMultipleThen[*MockInheritMultipleUseStdTypeWhen]{ then: _this.then, t: _this.t},
	}
}

// Defines the behavior when Close of the mock is called.
//
// As a default the method is expected to be called once.
// To change this behavior use the Times() method to define how often the function shall be called.
func (_this *MockInheritMultipleWhen) Close() *MockInheritMultipleCloseWhenWithTimes {
	for _, f := range _this.m.vClose {
		if f.validateArgs == nil {
			_this.m.t.Helper()
			_this.m.t.Fatalf("Unreachable condition. Call to 'Close' is already captured by previous WHEN statement.")
		}
	}
	var defaultExpected struct {
		fun func() (r0 error)
		expectedCalled int
		called int
		mutex sync.Mutex
	}
	defaultExpected.fun = func() (r0 error) { return }
	defaultExpected.expectedCalled = 1
	
	var validator struct {
		validateArgs func() bool
		expected []*struct {
			fun func() (r0 error)
			expectedCalled int
			called int
			mutex sync.Mutex
		}
	}
	validator.expected = append(validator.expected, &defaultExpected)
	_this.m.vClose = append(_this.m.vClose, &validator)
	var _then func() *MockInheritMultipleCloseWhen
	_then = func() *MockInheritMultipleCloseWhen {
		var _newExpected struct {
			fun func() (r0 error)
			expectedCalled int
			called int
			mutex sync.Mutex
		}
		_newExpected.fun = func() (r0 error) { return }
		_newExpected.expectedCalled = 1
		
		validator.expected = append(validator.expected, &_newExpected)
		return &MockInheritMultipleCloseWhen {
			expected: validator.expected,
			then: _then,
			t: _this.m.t,
		}
	}
	
	times := &MockInheritMultipleTimes[*MockInheritMultipleCloseWhen] {
		expectedCalled: &validator.expected[0].expectedCalled,
		then: _then,
		t: _this.m.t,
		MockInheritMultipleThen: MockInheritMultipleThen[*MockInheritMultipleCloseWhen]{ then: _then, t: _this.m.t},
	}
	return &MockInheritMultipleCloseWhenWithTimes {
		MockInheritMultipleCloseWhen: &MockInheritMultipleCloseWhen {
			expected: validator.expected,
			then: _then,
			t: _this.m.t,
		},
		MockInheritMultipleTimes: times,
	}
}

type MockInheritMultipleCloseWhen struct {
	expected []*struct {
		fun func() (r0 error)
		expectedCalled int
		called int
		mutex sync.Mutex
	}
	then func() *MockInheritMultipleCloseWhen
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
}

type MockInheritMultipleCloseWhenWithTimes struct {
	*MockInheritMultipleTimes[*MockInheritMultipleCloseWhen]
	*MockInheritMultipleCloseWhen
}

// Return the provided values when called
func (_this *MockInheritMultipleCloseWhen) Return(r0 error) *MockInheritMultipleTimes[*MockInheritMultipleCloseWhen] {
	_this.expected[len(_this.expected) -1].fun = func() (error) { return r0 }
	return &MockInheritMultipleTimes[*MockInheritMultipleCloseWhen] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockInheritMultipleThen: MockInheritMultipleThen[*MockInheritMultipleCloseWhen]{ then: _this.then, t: _this.t},
	}
}

// Do will execute the provided function and return the result when called
func (_this *MockInheritMultipleCloseWhen) Do(do func() (r0 error)) *MockInheritMultipleTimes[*MockInheritMultipleCloseWhen] {
	_this.expected[len(_this.expected) -1].fun = do
	return &MockInheritMultipleTimes[*MockInheritMultipleCloseWhen] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockInheritMultipleThen: MockInheritMultipleThen[*MockInheritMultipleCloseWhen]{ then: _this.then, t: _this.t},
	}
}

// Defines the behavior when Read of the mock is called.
//
// As a default the method is expected to be called once.
// To change this behavior use the Times() method to define how often the function shall be called.
func (_this *MockInheritMultipleWhen) Read() *MockInheritMultipleReadExpectWithTimes {
	for _, f := range _this.m.vRead {
		if f.validateArgs == nil {
			_this.m.t.Helper()
			_this.m.t.Fatalf("Unreachable condition. Call to 'Read' is already captured by previous WHEN statement.")
		}
	}
	var defaultExpected struct {
		fun func(p []byte) (n int, err error)
		expectedCalled int
		called int
		mutex sync.Mutex
	}
	defaultExpected.fun = func(p []byte) (n int, err error) { return }
	defaultExpected.expectedCalled = 1
	
	var validator struct {
		validateArgs func(p []byte) bool
		expected []*struct {
			fun func(p []byte) (n int, err error)
			expectedCalled int
			called int
			mutex sync.Mutex
		}
	}
	validator.expected = append(validator.expected, &defaultExpected)
	_this.m.vRead = append(_this.m.vRead, &validator)
	var _then func() *MockInheritMultipleReadWhen
	_then = func() *MockInheritMultipleReadWhen {
		var _newExpected struct {
			fun func(p []byte) (n int, err error)
			expectedCalled int
			called int
			mutex sync.Mutex
		}
		_newExpected.fun = func(p []byte) (n int, err error) { return }
		_newExpected.expectedCalled = 1
		
		validator.expected = append(validator.expected, &_newExpected)
		return &MockInheritMultipleReadWhen {
			expected: validator.expected,
			then: _then,
			t: _this.m.t,
		}
	}
	
	times := &MockInheritMultipleTimes[*MockInheritMultipleReadWhen] {
		expectedCalled: &validator.expected[0].expectedCalled,
		then: _then,
		t: _this.m.t,
		MockInheritMultipleThen: MockInheritMultipleThen[*MockInheritMultipleReadWhen]{ then: _then, t: _this.m.t},
	}
	return &MockInheritMultipleReadExpectWithTimes {
		MockInheritMultipleReadExpect: &MockInheritMultipleReadExpect {
			MockInheritMultipleReadWhen: &MockInheritMultipleReadWhen {
				expected: validator.expected,
				then: _then,
				t: _this.m.t,
			},
			validateArgs: &validator.validateArgs,
			times: times,
		},
		MockInheritMultipleTimes: times,
	}
}

type MockInheritMultipleReadExpect struct {
	*MockInheritMultipleReadWhen
	validateArgs *func(p []byte) bool
	times *MockInheritMultipleTimes[*MockInheritMultipleReadWhen]
}

// Expect will filter for given arguments.
// Each argument is matched with a filter function. Only if all arguments match this mocked function will be called.

// Arguments are either evaluated using the function, or ignored and always true if the function is set to nil.
func (_this *MockInheritMultipleReadExpect) Expect(p func([]byte) bool) *MockInheritMultipleReadWhenWithTimes {
	if !(p == nil) {
		*_this.validateArgs = func(_p []byte) bool {
			return (p == nil || p(_p))
		}
	}
	return &MockInheritMultipleReadWhenWithTimes {
		MockInheritMultipleReadWhen: _this.MockInheritMultipleReadWhen,
		MockInheritMultipleTimes: _this.times,
	}
}

type MockInheritMultipleReadExpectWithTimes struct {
	*MockInheritMultipleTimes[*MockInheritMultipleReadWhen]
	*MockInheritMultipleReadExpect
}

type MockInheritMultipleReadWhen struct {
	expected []*struct {
		fun func(p []byte) (n int, err error)
		expectedCalled int
		called int
		mutex sync.Mutex
	}
	then func() *MockInheritMultipleReadWhen
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
}

type MockInheritMultipleReadWhenWithTimes struct {
	*MockInheritMultipleTimes[*MockInheritMultipleReadWhen]
	*MockInheritMultipleReadWhen
}

// Return the provided values when called
func (_this *MockInheritMultipleReadWhen) Return(n int, err error) *MockInheritMultipleTimes[*MockInheritMultipleReadWhen] {
	_this.expected[len(_this.expected) -1].fun = func([]byte) (int, error) { return n, err }
	return &MockInheritMultipleTimes[*MockInheritMultipleReadWhen] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockInheritMultipleThen: MockInheritMultipleThen[*MockInheritMultipleReadWhen]{ then: _this.then, t: _this.t},
	}
}

// Do will execute the provided function and return the result when called
func (_this *MockInheritMultipleReadWhen) Do(do func(p []byte) (n int, err error)) *MockInheritMultipleTimes[*MockInheritMultipleReadWhen] {
	_this.expected[len(_this.expected) -1].fun = do
	return &MockInheritMultipleTimes[*MockInheritMultipleReadWhen] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockInheritMultipleThen: MockInheritMultipleThen[*MockInheritMultipleReadWhen]{ then: _this.then, t: _this.t},
	}
}

// Defines the behavior when Seek of the mock is called.
//
// As a default the method is expected to be called once.
// To change this behavior use the Times() method to define how often the function shall be called.
func (_this *MockInheritMultipleWhen) Seek() *MockInheritMultipleSeekExpectWithTimes {
	for _, f := range _this.m.vSeek {
		if f.validateArgs == nil {
			_this.m.t.Helper()
			_this.m.t.Fatalf("Unreachable condition. Call to 'Seek' is already captured by previous WHEN statement.")
		}
	}
	var defaultExpected struct {
		fun func(offset int64, whence int) (r0 int64, r1 error)
		expectedCalled int
		called int
		mutex sync.Mutex
	}
	defaultExpected.fun = func(offset int64, whence int) (r0 int64, r1 error) { return }
	defaultExpected.expectedCalled = 1
	
	var validator struct {
		validateArgs func(offset int64, whence int) bool
		expected []*struct {
			fun func(offset int64, whence int) (r0 int64, r1 error)
			expectedCalled int
			called int
			mutex sync.Mutex
		}
	}
	validator.expected = append(validator.expected, &defaultExpected)
	_this.m.vSeek = append(_this.m.vSeek, &validator)
	var _then func() *MockInheritMultipleSeekWhen
	_then = func() *MockInheritMultipleSeekWhen {
		var _newExpected struct {
			fun func(offset int64, whence int) (r0 int64, r1 error)
			expectedCalled int
			called int
			mutex sync.Mutex
		}
		_newExpected.fun = func(offset int64, whence int) (r0 int64, r1 error) { return }
		_newExpected.expectedCalled = 1
		
		validator.expected = append(validator.expected, &_newExpected)
		return &MockInheritMultipleSeekWhen {
			expected: validator.expected,
			then: _then,
			t: _this.m.t,
		}
	}
	
	times := &MockInheritMultipleTimes[*MockInheritMultipleSeekWhen] {
		expectedCalled: &validator.expected[0].expectedCalled,
		then: _then,
		t: _this.m.t,
		MockInheritMultipleThen: MockInheritMultipleThen[*MockInheritMultipleSeekWhen]{ then: _then, t: _this.m.t},
	}
	return &MockInheritMultipleSeekExpectWithTimes {
		MockInheritMultipleSeekExpect: &MockInheritMultipleSeekExpect {
			MockInheritMultipleSeekWhen: &MockInheritMultipleSeekWhen {
				expected: validator.expected,
				then: _then,
				t: _this.m.t,
			},
			validateArgs: &validator.validateArgs,
			times: times,
		},
		MockInheritMultipleTimes: times,
	}
}

type MockInheritMultipleSeekExpect struct {
	*MockInheritMultipleSeekWhen
	validateArgs *func(offset int64, whence int) bool
	times *MockInheritMultipleTimes[*MockInheritMultipleSeekWhen]
}

// Expect will filter for given arguments.
// Each argument is matched with a filter function. Only if all arguments match this mocked function will be called.

// Arguments are either evaluated using the function, or ignored and always true if the function is set to nil.
func (_this *MockInheritMultipleSeekExpect) Expect(offset func(int64) bool, whence func(int) bool) *MockInheritMultipleSeekWhenWithTimes {
	if !(offset == nil && whence == nil) {
		*_this.validateArgs = func(_offset int64, _whence int) bool {
			return (offset == nil || offset(_offset)) && (whence == nil || whence(_whence))
		}
	}
	return &MockInheritMultipleSeekWhenWithTimes {
		MockInheritMultipleSeekWhen: _this.MockInheritMultipleSeekWhen,
		MockInheritMultipleTimes: _this.times,
	}
}

type MockInheritMultipleSeekExpectWithTimes struct {
	*MockInheritMultipleTimes[*MockInheritMultipleSeekWhen]
	*MockInheritMultipleSeekExpect
}

type MockInheritMultipleSeekWhen struct {
	expected []*struct {
		fun func(offset int64, whence int) (r0 int64, r1 error)
		expectedCalled int
		called int
		mutex sync.Mutex
	}
	then func() *MockInheritMultipleSeekWhen
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
}

type MockInheritMultipleSeekWhenWithTimes struct {
	*MockInheritMultipleTimes[*MockInheritMultipleSeekWhen]
	*MockInheritMultipleSeekWhen
}

// Return the provided values when called
func (_this *MockInheritMultipleSeekWhen) Return(r0 int64, r1 error) *MockInheritMultipleTimes[*MockInheritMultipleSeekWhen] {
	_this.expected[len(_this.expected) -1].fun = func(int64, int) (int64, error) { return r0, r1 }
	return &MockInheritMultipleTimes[*MockInheritMultipleSeekWhen] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockInheritMultipleThen: MockInheritMultipleThen[*MockInheritMultipleSeekWhen]{ then: _this.then, t: _this.t},
	}
}

// Do will execute the provided function and return the result when called
func (_this *MockInheritMultipleSeekWhen) Do(do func(offset int64, whence int) (r0 int64, r1 error)) *MockInheritMultipleTimes[*MockInheritMultipleSeekWhen] {
	_this.expected[len(_this.expected) -1].fun = do
	return &MockInheritMultipleTimes[*MockInheritMultipleSeekWhen] {
		expectedCalled: &_this.expected[len(_this.expected) -1].expectedCalled,
		then: _this.then,
		t: _this.t,
		MockInheritMultipleThen: MockInheritMultipleThen[*MockInheritMultipleSeekWhen]{ then: _this.then, t: _this.t},
	}
}

type MockInheritMultipleThen [T any] struct {
	then func() T
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
}

// Then continue with another action
func (_this *MockInheritMultipleThen[T]) Then() T {
	_this.t.Helper()
	return _this.then()
}

type MockInheritMultipleTimes[T any] struct {
	expectedCalled *int
	then func() T
	t interface {
		Fatalf(format string, args ...any)
		Helper()
	}
	MockInheritMultipleThen[T]
}

// Times sets how often the mocked function is expected to be called.
// Test will fail if the number of calls do not match with the expected calls value.
func (_this *MockInheritMultipleTimes[T]) Times(times int) *MockInheritMultipleThen[T] {
	_this.t.Helper()
	*_this.expectedCalled = times
	retVal := &MockInheritMultipleThen[T] { t: _this.t, then: _this.then }
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
func (_this *MockInheritMultipleTimes[T]) AnyTimes() {
	*_this.expectedCalled = -1
}

// Never will fail if the function is ever called.
func (_this *MockInheritMultipleTimes[T]) Never() {
	*_this.expectedCalled = 0
}

