package model_test

import (
	"testing"

	"github.com/becheran/smock/model"
	"github.com/stretchr/testify/assert"
)

func TestImportName(t *testing.T) {
	assert.Equal(t, "", model.Import{}.ImportName())
	assert.Equal(t, "", model.Import{Path: "foo/bar/"}.ImportName())
	assert.Equal(t, "", model.Import{Path: "/"}.ImportName())
	assert.Equal(t, "bar", model.Import{Name: "bar"}.ImportName())
	assert.Equal(t, "bar", model.Import{Name: "bar", Path: "foo/other"}.ImportName())
	assert.Equal(t, "bar", model.Import{Path: "foo/bar"}.ImportName())
	assert.Equal(t, "bar", model.Import{Path: "bar"}.ImportName())
}

func TestImportString(t *testing.T) {
	assert.Equal(t, ` ""`, model.Import{}.String())
	assert.Equal(t, `other "foo/other"`, model.Import{Path: "foo/other"}.String())
	assert.Equal(t, `rename "foo/other"`, model.Import{Name: "rename", Path: "foo/other"}.String())
}
