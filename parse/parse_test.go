package parse_test

import (
	"fmt"
	"go/parser"
	"go/token"
	"testing"

	"github.com/becheran/smock/model"
	"github.com/becheran/smock/parse"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseInterface(t *testing.T) {
	const (
		src1 = `package p

var X = f(3.14)*2 + c

// Comment
type MyInterface interface {
	other.Inter
	Foo(x other.Type, bar, baz string) (o other.Other, oo map[string]SamePackage)
	unexported() int
}
`

		src2 = `package p

type MyInterface other.Other`
	)

	var suite = []struct {
		src         string
		line        int
		errContains string
		res         model.InterfaceResult
	}{
		{src1, 1, "unexpected identifier", model.InterfaceResult{}},
		{src1, 18, "interface not found", model.InterfaceResult{}},

		{src1, 5, "", model.InterfaceResult{
			Name:       "MyInterface",
			References: []model.Reference{{PackageID: "other", Name: "Inter"}},
			Methods: []model.Method{{
				Name:    "Foo",
				Params:  []model.Ident{{Name: "x", Type: "other.Type"}, {Name: "bar", Type: "string"}, {Name: "baz", Type: "string"}},
				Results: []model.Ident{{Name: "o", Type: "other.Other"}, {Name: "oo", Type: "map[string]p.SamePackage"}},
			}},
		}},

		{src2, 2, "", model.InterfaceResult{
			Name:       "MyInterface",
			References: []model.Reference{{PackageID: "other", Name: "Other"}},
		}},
	}
	for idx, s := range suite {
		t.Run(fmt.Sprintf("%d", idx), func(t *testing.T) {
			fset := token.NewFileSet()
			f, err := parser.ParseFile(fset, "src.go", s.src, 0)
			require.Nil(t, err)

			res, err := parse.ParseInterface(fset, f, s.line)
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
