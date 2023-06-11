package model_test

import (
	"fmt"
	"testing"

	"github.com/becheran/smock/model"
	"github.com/stretchr/testify/assert"
)

func TestMethodSignature(t *testing.T) {
	assert.Equal(t, "()", model.Method{}.Signature())
	assert.Equal(t, "() (foo string)", model.Method{
		Results: model.IdentList{{Name: "foo", Type: "string"}},
	}.Signature())
	assert.Equal(t, fmt.Sprintf("(%s0 int, %s1 bool) (%s0 int)", model.IdentTypeInput, model.IdentTypeInput, model.IdentTypeResult), model.Method{
		Params:  model.IdentList{{Type: "int"}, {Type: "bool"}},
		Results: model.IdentList{{Type: "int"}},
	}.Signature())
}

func TestImportName(t *testing.T) {
	assert.Equal(t, "", model.Import{}.ImportName())
	assert.Equal(t, "", model.Import{Path: "foo/bar/"}.ImportName())
	assert.Equal(t, "", model.Import{Path: "/"}.ImportName())
	assert.Equal(t, "bar", model.Import{Name: "bar"}.ImportName())
	assert.Equal(t, "bar", model.Import{Name: "bar", Path: "foo/other"}.ImportName())
	assert.Equal(t, "bar", model.Import{Path: "foo/bar"}.ImportName())
	assert.Equal(t, "bar", model.Import{Path: "bar"}.ImportName())
	assert.Equal(t, "bar", model.Import{Path: "foo/go-bar"}.ImportName())
	assert.Equal(t, "bar", model.Import{Path: "foo/foo-bar"}.ImportName())
}

func TestImportString(t *testing.T) {
	assert.Equal(t, ` ""`, model.Import{}.String())
	assert.Equal(t, `other "foo/other"`, model.Import{Path: "foo/other"}.String())
	assert.Equal(t, `rename "foo/other"`, model.Import{Name: "rename", Path: "foo/other"}.String())
}
