// DO NOT EDIT
// Code generated by smock 

package testpackage_mock

import (
	testpackage "github.com/test/testpackage"
	"fmt"
)

// MockWithLambda must implement interface testpackage.WithLambda
var _ testpackage.WithLambda = &MockWithLambda{}

func NewMockWithLambda(t interface {
	Fatalf(format string, args ...interface{})
	Helper()
}) *MockWithLambda {
	return &MockWithLambda{t: t}
}

type MockWithLambda struct {
	t interface {
		Fatalf(format string, args ...interface{})
		Helper()
	}
	
	fFoo func(a int, b ...string)
	fEmpty func()
}

func (m *MockWithLambda) Foo(a int, b ...string) {
	if m.fFoo != nil {
		m.fFoo(a, b...)
	} else {
		m.unexpectedCall("Foo", fmt.Sprintf("%+v, %+v", a, b))
		return
	}
}

func (m *MockWithLambda) Empty() {
	if m.fEmpty != nil {
		m.fEmpty()
	} else {
		m.unexpectedCall("Empty", fmt.Sprintf(""))
		return
	}
}

func (m *MockWithLambda) WHEN() *MockWithLambdaWhen {
	return &MockWithLambdaWhen{
		m: m,
	}
}

func (m *MockWithLambda) unexpectedCall(method, args string) {
	m.t.Helper()
	m.t.Fatalf(`Unexpected call to MockWithLambda.%s(%s)`, method, args)
}

type MockWithLambdaWhen struct {
	m *MockWithLambda
}

func (mh *MockWithLambdaWhen) Foo() *MockWithLambdaFooFunc {
	mh.m.fFoo = func(a int, b ...string) { return }
	return &MockWithLambdaFooFunc{m: mh.m}
}

type MockWithLambdaFooFunc struct {
	m *MockWithLambda
}

func (f *MockWithLambdaFooFunc) Do(do func(a int, b ...string)) {
	f.m.fFoo = do
}

func (mh *MockWithLambdaWhen) Empty() *MockWithLambdaEmptyFunc {
	mh.m.fEmpty = func() { return }
	return &MockWithLambdaEmptyFunc{m: mh.m}
}

type MockWithLambdaEmptyFunc struct {
	m *MockWithLambda
}

func (f *MockWithLambdaEmptyFunc) Do(do func()) {
	f.m.fEmpty = do
}
