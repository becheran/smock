package testpackage

import (
	"io"
	"os"
)

//go:generate smock -debug
type Simple interface {
	Foo()
	SingleArg(int)
	Bar(a int, b string, c struct{}, d *struct{}, e any, f []byte) string
	Baz(a int, b string) (s string)
	Fun(a, b func(a, b func(o, i string) (f, a int))) (r, r2 func())
}

//go:generate smock -debug
type Extend interface {
	RetType() MyType
	UseStdType(fi os.FileInfo) io.Reader
}

//go:generate smock -debug
type InheritOwn Extend

//go:generate smock -debug
type InheritExt io.ReadSeekCloser

//go:generate smock -debug
type InheritMultiple interface {
	io.ReadSeekCloser
	Extend
	Own(int, string) (int, string)
}

//go:generate smock -debug
type WithTypes[T, B any] interface {
	Foo(a, b T) B
	Empty()
}

//go:generate smock -debug
type WithLambda[T comparable] interface {
	Foo(a int, b ...string) bool
	Bar(b ...struct{}) bool
	Baz(b ...T) bool
}
