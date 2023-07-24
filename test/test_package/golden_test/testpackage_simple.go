// DO NOT EDIT
// Code generated by smock 

package testpackage_mock

import (
	testpackage "github.com/test/testpackage"
	"fmt"
)

// MockSimple must implement interface testpackage.Simple
var _ testpackage.Simple = &MockSimple{}

func NewMockSimple(t interface {
	Fatalf(format string, args ...interface{})
	Helper()
}) *MockSimple {
	return &MockSimple{t: t}
}

type MockSimple struct {
	t interface {
		Fatalf(format string, args ...interface{})
		Helper()
	}
	fFoo func()
	fBar func(a int, b string) (r0 string)
	fBaz func(a string, b string) (r int, r2 int)
	fFun func(a func(func(string, string) (int, int), func(string, string) (int, int)), b func(func(string, string) (int, int), func(string, string) (int, int))) (r func(), r2 func())
}

func (m *MockSimple) Foo() {
	if m.fFoo != nil {
		m.fFoo()
	} else {
		m.unexpectedCall("Foo", fmt.Sprintf(""))
		return
	}
}

func (m *MockSimple) Bar(a int, b string) (r0 string) {
	if m.fBar != nil {
		return m.fBar(a, b)
	} else {
		m.unexpectedCall("Bar", fmt.Sprintf("%+v, %+v", a, b))
		return
	}
}

func (m *MockSimple) Baz(a string, b string) (r int, r2 int) {
	if m.fBaz != nil {
		return m.fBaz(a, b)
	} else {
		m.unexpectedCall("Baz", fmt.Sprintf("%+v, %+v", a, b))
		return
	}
}

func (m *MockSimple) Fun(a func(func(string, string) (int, int), func(string, string) (int, int)), b func(func(string, string) (int, int), func(string, string) (int, int))) (r func(), r2 func()) {
	if m.fFun != nil {
		return m.fFun(a, b)
	} else {
		m.unexpectedCall("Fun", fmt.Sprintf("%+v, %+v", a, b))
		return
	}
}

func (m *MockSimple) WHEN() *MockSimpleWhen {
	return &MockSimpleWhen{
		m: m,
	}
}

func (m *MockSimple) unexpectedCall(method, args string) {
	m.t.Helper()
	m.t.Fatalf(`Unexpected call to MockSimple.%s(%s)`, method, args)
}

type MockSimpleWhen struct {
	m *MockSimple
}

func (mh *MockSimpleWhen) Foo() *MockSimpleFooFunc {
	mh.m.fFoo = func() { return }
	return &MockSimpleFooFunc{m: mh.m}
}

type MockSimpleFooFunc struct {
	m *MockSimple
}

func (f *MockSimpleFooFunc) Do(do func()) {
	f.m.fFoo = do
}

func (mh *MockSimpleWhen) Bar() *MockSimpleBarFunc {
	mh.m.fBar = func(a int, b string) (r0 string) { return }
	return &MockSimpleBarFunc{m: mh.m}
}

type MockSimpleBarFunc struct {
	m *MockSimple
}

func (f *MockSimpleBarFunc) Return(r0 string) {
	f.m.fBar = func(int, string) (string) { return r0 }
}

func (f *MockSimpleBarFunc) Do(do func(a int, b string) (r0 string)) {
	f.m.fBar = do
}

func (mh *MockSimpleWhen) Baz() *MockSimpleBazFunc {
	mh.m.fBaz = func(a string, b string) (r int, r2 int) { return }
	return &MockSimpleBazFunc{m: mh.m}
}

type MockSimpleBazFunc struct {
	m *MockSimple
}

func (f *MockSimpleBazFunc) Return(r int, r2 int) {
	f.m.fBaz = func(string, string) (int, int) { return r, r2 }
}

func (f *MockSimpleBazFunc) Do(do func(a string, b string) (r int, r2 int)) {
	f.m.fBaz = do
}

func (mh *MockSimpleWhen) Fun() *MockSimpleFunFunc {
	mh.m.fFun = func(a func(func(string, string) (int, int), func(string, string) (int, int)), b func(func(string, string) (int, int), func(string, string) (int, int))) (r func(), r2 func()) { return }
	return &MockSimpleFunFunc{m: mh.m}
}

type MockSimpleFunFunc struct {
	m *MockSimple
}

func (f *MockSimpleFunFunc) Return(r func(), r2 func()) {
	f.m.fFun = func(func(func(string, string) (int, int), func(string, string) (int, int)), func(func(string, string) (int, int), func(string, string) (int, int))) (func(), func()) { return r, r2 }
}

func (f *MockSimpleFunFunc) Do(do func(a func(func(string, string) (int, int), func(string, string) (int, int)), b func(func(string, string) (int, int), func(string, string) (int, int))) (r func(), r2 func())) {
	f.m.fFun = do
}
