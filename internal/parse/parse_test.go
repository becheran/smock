package parse_test

import (
	"fmt"
	"go/parser"
	"go/token"
	"testing"

	"github.com/becheran/smock/internal/model"
	"github.com/becheran/smock/internal/parse"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	src1 = `package p

import (
	other "github.com/foo/bar"
	"io"
)

var X = f(3.14)*2 + c

// Comment
type MyInterface interface {
	Foo(x other.Type, bar, baz *string, r io.Reader, fun func(), funArg func(a, b string, c int) (r0, r1 int)) (o other.Other, oo map[string]SamePackage)
	Void()
}
`

	src2 = `package p

import (
	"github.com/foo/bar/other"
	"golang.org/x/exp/constraints"
)

type MyInterface other.Other

type InterfaceWithType[T any, B constraints.Ordered] interface {
	Foo(a, b B) T
}

type RefMy MyInterface

type Unexported interface {
	unexported()
}
`
)

func TestParseInterface(t *testing.T) {
	var suite = []struct {
		src         string
		line        int
		errContains string
		res         model.InterfaceResult
	}{
		{src1, 1, "unexpected identifier", model.InterfaceResult{}},
		{src1, 18, "interface at p:18 not found", model.InterfaceResult{}},

		{src1, 8, "", model.InterfaceResult{
			Name:        "MyInterface",
			PackageName: "p",
			Methods: []model.Method{
				{
					Name: "Foo",
					Params: []model.Ident{
						{Name: "x", Type: "other.Type"},
						{Name: "bar", Type: "*string"},
						{Name: "baz", Type: "*string"},
						{Name: "r", Type: "io.Reader"},
						{Name: "fun", Type: "func()"},
						{Name: "funArg", Type: "func(string, string, int) (int, int)"},
					},
					Results: []model.Ident{{Name: "o", Type: "other.Other"}, {Name: "oo", Type: "map[string]p.SamePackage"}}},
				{Name: "Void", Params: nil, Results: nil},
			},
			Imports: []model.Import{
				{Path: "io"},
				{Name: "other", Path: "github.com/foo/bar"},
			}}},

		{src2, 8, "", model.InterfaceResult{
			Name:        "InterfaceWithType",
			PackageName: "p",
			Methods: []model.Method{{
				Name:    "Foo",
				Params:  []model.Ident{{Name: "a", Type: "B"}, {Name: "b", Type: "B"}},
				Results: []model.Ident{{Type: "T"}},
			}},
			Types: []model.Ident{{Name: "T", Type: "any"}, {Name: "B", Type: "constraints.Ordered"}},
			Imports: []model.Import{
				{Name: "", Path: "golang.org/x/exp/constraints"},
			}}},
		{src2, 6, "failed to resolve package reference", model.InterfaceResult{}},
		{src2, 13, "interface 'MyInterface' not found", model.InterfaceResult{}},
		{src2, 15, "unexported", model.InterfaceResult{}},
	}
	for idx, s := range suite {
		t.Run(fmt.Sprintf("%d", idx), func(t *testing.T) {
			fset := token.NewFileSet()
			f, err := parser.ParseFile(fset, "src.go", s.src, 0)
			require.Nil(t, err)

			res, err := parse.ParseInterfaceAtPosition(fset, f, s.line)
			if s.errContains != "" {
				assert.ErrorContains(t, err, s.errContains)
				assert.Empty(t, res)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, s.res, res)
			}
		})
	}
}
