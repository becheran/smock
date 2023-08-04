// DO NOT EDIT
// Code generated by smock 

package testpackage_mock

import (
	"fmt"
	"reflect"
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
	
	vFoo []*struct{fun func(a T, b T) (r0 B); validateArgs func(a T, b T) bool}
	vEmpty []*struct{fun func(); validateArgs func() bool}
}

func (_this *MockWithTypes[T, B]) Foo(a T, b T) (r0 B) {
	for _, _check := range _this.vFoo {
		if _check.validateArgs == nil || _check.validateArgs(a, b) {
			return _check.fun(a, b)
		}
	}
	_this.t.Helper()
	_this.unexpectedCall("Foo", a, b)
	return
}

func (_this *MockWithTypes[T, B]) Empty() {
	for _, _check := range _this.vEmpty {
		if _check.validateArgs == nil || _check.validateArgs() {
			_check.fun()
			return
		}
	}
	_this.t.Helper()
	_this.unexpectedCall("Empty", )
}

func (_this *MockWithTypes[T, B]) unexpectedCall(method string, args ...any) {
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
	_this.t.Fatalf(`Unexpected call %s(%s)`, method, argsStr)
}

func (_this *MockWithTypes[T, B]) WHEN() *MockWithTypesWhen[T, B] {
	return &MockWithTypesWhen[T, B]{
		m: _this,
	}
}

type MockWithTypesWhen[T any, B any] struct {
	m *MockWithTypes[T, B]
}

func (_this *MockWithTypesWhen[T, B]) Foo() *MockWithTypesFooArgs[T, B] {
	for _, f := range _this.m.vFoo {
		if f.validateArgs == nil {
			_this.m.t.Helper()
			_this.m.t.Fatalf("Unreachable condition. Call to 'Foo' is already captured by previous WHEN statement.")
		}
	}
	var validator struct {
		fun func(a T, b T) (r0 B)
		validateArgs func(a T, b T) bool
	}
	validator.fun = func(a T, b T) (r0 B) { return }
	_this.m.vFoo = append(_this.m.vFoo, &validator)
	return &MockWithTypesFooArgs[T, B] {
		MockWithTypesFooArgsEval: MockWithTypesFooArgsEval[T, B]{fun: &validator.fun},
		validateArgs: &validator.validateArgs,
		fun: &validator.fun,
	}
}

type MockWithTypesFooArgs[T any, B any] struct {
	MockWithTypesFooArgsEval[T, B]
	fun *func(a T, b T) (r0 B)
	validateArgs *func(a T, b T) bool
}

func (_this *MockWithTypesFooArgs[T, B]) Expect(a func(T) bool, b func(T) bool) *MockWithTypesFooArgsEval[T, B] {
	if !(a == nil && b == nil) {
		*_this.validateArgs = func(_a T, _b T) bool {
			return (a == nil || a(_a)) && (b == nil || b(_b))
		}
	}
	return &_this.MockWithTypesFooArgsEval
}

type MockWithTypesFooArgsEval[T any, B any] struct {
	fun *func(a T, b T) (r0 B)
}

func (_this *MockWithTypesFooArgsEval[T, B]) Return(r0 B) {
	*_this.fun = func(T, T) (B) { return r0 }
}

func (_this *MockWithTypesFooArgsEval[T, B]) Do(do func(a T, b T) (r0 B)) {
	*_this.fun = do
}

func (_this *MockWithTypesWhen[T, B]) Empty() *MockWithTypesEmptyArgsEval[T, B] {
	for _, f := range _this.m.vEmpty {
		if f.validateArgs == nil {
			_this.m.t.Helper()
			_this.m.t.Fatalf("Unreachable condition. Call to 'Empty' is already captured by previous WHEN statement.")
		}
	}
	var validator struct {
		fun func()
		validateArgs func() bool
	}
	validator.fun = func() { }
	_this.m.vEmpty = append(_this.m.vEmpty, &validator)
	return &MockWithTypesEmptyArgsEval[T, B] {
		fun: &validator.fun,
	}
}

type MockWithTypesEmptyArgsEval[T any, B any] struct {
	fun *func()
}

func (_this *MockWithTypesEmptyArgsEval[T, B]) Do(do func()) {
	*_this.fun = do
}
