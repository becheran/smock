// DO NOT EDIT
// Code generated by smock 

package testpackage_mock

import (
	"fmt"
)

func NewMockWithTypes[T any, B any](t interface {
	Fatalf(format string, args ...interface{})
	Helper()
}) *MockWithTypes[T, B] {
	return &MockWithTypes[T, B]{t: t}
}

type MockWithTypes[T any, B any] struct {
	t interface {
		Fatalf(format string, args ...interface{})
		Helper()
	}
	fFoo func(a T, b T) (r0 B)
	fEmpty func()
}

func (m *MockWithTypes[T, B]) Foo(a T, b T) (r0 B) {
	if m.fFoo != nil {
		return m.fFoo(a, b)
	} else {
		m.unexpectedCall("Foo", fmt.Sprintf("%+v, %+v", a, b))
		return
	}
}

func (m *MockWithTypes[T, B]) Empty() {
	if m.fEmpty != nil {
		m.fEmpty()
	} else {
		m.unexpectedCall("Empty", fmt.Sprintf(""))
		return
	}
}

func (m *MockWithTypes[T, B]) WHEN() *MockWithTypesWhen[T, B] {
	return &MockWithTypesWhen[T, B]{
		m: m,
	}
}

func (m *MockWithTypes[T, B]) unexpectedCall(method, args string) {
	m.t.Helper()
	m.t.Fatalf(`Unexpected call to MockWithTypes.%s(%s)`, method, args)
}

type MockWithTypesWhen[T any, B any] struct {
	m *MockWithTypes[T, B]
}

func (mh *MockWithTypesWhen[T, B]) Foo() *MockWithTypesFooFunc[T, B] {
	mh.m.fFoo = func(a T, b T) (r0 B) { return }
	return &MockWithTypesFooFunc[T, B]{m: mh.m}
}

type MockWithTypesFooFunc[T any, B any] struct {
	m *MockWithTypes[T, B]
}

func (f *MockWithTypesFooFunc[T, B]) Return(r0 B) {
	f.m.fFoo = func(T, T) (B) { return r0 }
}

func (f *MockWithTypesFooFunc[T, B]) Do(do func(a T, b T) (r0 B)) {
	f.m.fFoo = do
}

func (mh *MockWithTypesWhen[T, B]) Empty() *MockWithTypesEmptyFunc[T, B] {
	mh.m.fEmpty = func() { return }
	return &MockWithTypesEmptyFunc[T, B]{m: mh.m}
}

type MockWithTypesEmptyFunc[T any, B any] struct {
	m *MockWithTypes[T, B]
}

func (f *MockWithTypesEmptyFunc[T, B]) Do(do func()) {
	f.m.fEmpty = do
}
