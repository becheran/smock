// DO NOT EDIT
// Code generated by smock 

package testpackage_mock

import (
	testpackage "github.com/test/testpackage"
	"fmt"
	"reflect"
)

// MockInheritExt must implement interface testpackage.InheritExt
var _ testpackage.InheritExt = &MockInheritExt{}

func NewMockInheritExt(t interface {
	Fatalf(format string, args ...interface{})
	Helper()
}) *MockInheritExt {
	return &MockInheritExt{t: t}
}

type MockInheritExt struct {
	t interface {
		Fatalf(format string, args ...interface{})
		Helper()
	}
	
	vClose []*struct{fun func() (r0 error); validateArgs func() bool}
	vRead []*struct{fun func(p []byte) (n int, err error); validateArgs func(p []byte) bool}
	vSeek []*struct{fun func(offset int64, whence int) (r0 int64, r1 error); validateArgs func(offset int64, whence int) bool}
}

func (m *MockInheritExt) Close() (r0 error) {
	for _, check := range m.vClose {
		if check.validateArgs == nil || check.validateArgs() {
			return check.fun()
		}
	}
	m.unexpectedCall("Close", )
	return
}

func (m *MockInheritExt) Read(p []byte) (n int, err error) {
	for _, check := range m.vRead {
		if check.validateArgs == nil || check.validateArgs(p) {
			return check.fun(p)
		}
	}
	m.unexpectedCall("Read", p)
	return
}

func (m *MockInheritExt) Seek(offset int64, whence int) (r0 int64, r1 error) {
	for _, check := range m.vSeek {
		if check.validateArgs == nil || check.validateArgs(offset, whence) {
			return check.fun(offset, whence)
		}
	}
	m.unexpectedCall("Seek", offset, whence)
	return
}

func (m *MockInheritExt) unexpectedCall(method string, args ...any) {
	argsStr := ""
	for idx, arg := range args {
		t := reflect.TypeOf(arg)
		if t.Kind() == reflect.Func {
			argsStr += fmt.Sprintf("%T", t)
		} else {
			argsStr += fmt.Sprintf("%+v", t)
		}
		if idx+1 < len(args) {
			argsStr += ", "
		}
	}
	m.t.Helper()
	m.t.Fatalf(`Unexpected call to MockInheritExt.%s(%s)`, method, argsStr)
}

func (m *MockInheritExt) WHEN() *MockInheritExtWhen {
	return &MockInheritExtWhen{
		m: m,
	}
}

type MockInheritExtWhen struct {
	m *MockInheritExt
}

func (mh *MockInheritExtWhen) Close() *MockInheritExtCloseArgsEval {
	for _, f := range  mh.m.vClose {
		if f.validateArgs == nil {
			mh.m.t.Helper()
			mh.m.t.Fatalf("Unreachable condition. Call to 'Close' is already captured by previous WHEN statement.")
		}
	}
	var validator struct {
		fun func() (r0 error)
		validateArgs func() bool
	}
	validator.fun = func() (r0 error) { return }
	mh.m.vClose = append(mh.m.vClose, &validator)
	return &MockInheritExtCloseArgsEval {
		fun: &validator.fun,
	}
}

type MockInheritExtCloseArgsEval struct {
	fun *func() (r0 error)
}

func (f *MockInheritExtCloseArgsEval) Return(r0 error) {
	*f.fun = func() (error) { return r0 }
}

func (f *MockInheritExtCloseArgsEval) Do(do func() (r0 error)) {
	*f.fun = do
}

func (mh *MockInheritExtWhen) Read() *MockInheritExtReadArgs {
	for _, f := range  mh.m.vRead {
		if f.validateArgs == nil {
			mh.m.t.Helper()
			mh.m.t.Fatalf("Unreachable condition. Call to 'Read' is already captured by previous WHEN statement.")
		}
	}
	var validator struct {
		fun func(p []byte) (n int, err error)
		validateArgs func(p []byte) bool
	}
	validator.fun = func(p []byte) (n int, err error) { return }
	mh.m.vRead = append(mh.m.vRead, &validator)
	return &MockInheritExtReadArgs {
		MockInheritExtReadArgsEval: MockInheritExtReadArgsEval{fun: &validator.fun},
		validateArgs: &validator.validateArgs,
		fun: &validator.fun,
	}
}

type MockInheritExtReadArgs struct {
	MockInheritExtReadArgsEval
	fun *func(p []byte) (n int, err error)
	validateArgs *func(p []byte) bool
}

func (f *MockInheritExtReadArgs) Expect(p func([]byte) bool) *MockInheritExtReadArgsEval {
	if !(p == nil) {
		*f.validateArgs = func(matchp []byte) bool {
			return (p == nil || p(matchp))
		}
	}
	return &f.MockInheritExtReadArgsEval
}

type MockInheritExtReadArgsEval struct {
	fun *func(p []byte) (n int, err error)
}

func (f *MockInheritExtReadArgsEval) Return(n int, err error) {
	*f.fun = func([]byte) (int, error) { return n, err }
}

func (f *MockInheritExtReadArgsEval) Do(do func(p []byte) (n int, err error)) {
	*f.fun = do
}

func (mh *MockInheritExtWhen) Seek() *MockInheritExtSeekArgs {
	for _, f := range  mh.m.vSeek {
		if f.validateArgs == nil {
			mh.m.t.Helper()
			mh.m.t.Fatalf("Unreachable condition. Call to 'Seek' is already captured by previous WHEN statement.")
		}
	}
	var validator struct {
		fun func(offset int64, whence int) (r0 int64, r1 error)
		validateArgs func(offset int64, whence int) bool
	}
	validator.fun = func(offset int64, whence int) (r0 int64, r1 error) { return }
	mh.m.vSeek = append(mh.m.vSeek, &validator)
	return &MockInheritExtSeekArgs {
		MockInheritExtSeekArgsEval: MockInheritExtSeekArgsEval{fun: &validator.fun},
		validateArgs: &validator.validateArgs,
		fun: &validator.fun,
	}
}

type MockInheritExtSeekArgs struct {
	MockInheritExtSeekArgsEval
	fun *func(offset int64, whence int) (r0 int64, r1 error)
	validateArgs *func(offset int64, whence int) bool
}

func (f *MockInheritExtSeekArgs) Expect(offset func(int64) bool, whence func(int) bool) *MockInheritExtSeekArgsEval {
	if !(offset == nil && whence == nil) {
		*f.validateArgs = func(matchoffset int64, matchwhence int) bool {
			return (offset == nil || offset(matchoffset)) && (whence == nil || whence(matchwhence))
		}
	}
	return &f.MockInheritExtSeekArgsEval
}

type MockInheritExtSeekArgsEval struct {
	fun *func(offset int64, whence int) (r0 int64, r1 error)
}

func (f *MockInheritExtSeekArgsEval) Return(r0 int64, r1 error) {
	*f.fun = func(int64, int) (int64, error) { return r0, r1 }
}

func (f *MockInheritExtSeekArgsEval) Do(do func(offset int64, whence int) (r0 int64, r1 error)) {
	*f.fun = do
}
