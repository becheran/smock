package testpackage

import (
	"io"
	"net/url"
	"os"

	renamed "github.com/test/testpackage/other"
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
	renamed.Bar
	io.ReadSeekCloser
	Extend
	Own(int, string) (int, string)
	Read(p []byte) (n int, err error) // Override io.ReaderSeekerCloser
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

//go:generate smock -debug
type unexported interface {
	Foo()
}

//go:generate smock -debug
type UseUrlWithUrlName interface {
	InUrl(url url.URL, fun func(url url.URL))
	RetUrl() (url url.URL, fun func() url.URL)
}
