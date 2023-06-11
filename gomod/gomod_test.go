package gomod_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/becheran/smock/gomod"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindModNoFile(t *testing.T) {
	root := t.TempDir()
	path := root + "/foo/bar/"
	*gomod.RootDir = root

	require.Nil(t, os.MkdirAll(path, os.ModePerm))

	info, err := gomod.FindMod(path)
	assert.Empty(t, info)
	assert.ErrorContains(t, err, "not found")
}

func TestFindModInvalid(t *testing.T) {
	root := t.TempDir()
	path := root + "/foo/bar/"
	*gomod.RootDir = root

	require.Nil(t, os.MkdirAll(path, os.ModePerm))
	require.Nil(t, os.WriteFile(root+"/foo/go.mod", []byte{}, os.ModePerm))

	info, err := gomod.FindMod(path)
	assert.Empty(t, info)
	assert.ErrorContains(t, err, "failed to parse")
}

func TestFindMod(t *testing.T) {
	root := t.TempDir()
	p := root + "/foo/bar/myfile.go"
	*gomod.RootDir = root

	require.Nil(t, os.MkdirAll(p, os.ModePerm))
	require.Nil(t, os.WriteFile(root+"/gO.mod", []byte(`module github.com/becheran/smock

go 1.20

require github.com/stretchr/testify v1.8.2

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
`), os.ModePerm))

	info, err := gomod.FindMod(p)
	assert.Nil(t, err)
	assert.Equal(t, "github.com/becheran/smock", info.ModuleName)
	assert.Equal(t, gomod.PathToUnix(root), info.Path)
}

func TestModImportPath(t *testing.T) {
	var suite = []struct {
		info   gomod.ModInfo
		path   string
		result string
	}{
		{gomod.ModInfo{ModuleName: "github.com/becheran/smock", Path: "/foo/bar"}, "/foo/bar", "github.com/becheran/smock"},
		{gomod.ModInfo{ModuleName: "github.com/becheran/smock", Path: "/foo/bar"}, "/foo/bar/baz", "github.com/becheran/smock/baz"},
		{gomod.ModInfo{ModuleName: "github.com/becheran/smock", Path: "C:\\foo\\bar\\smock"}, "C:\\foo\\bar\\smock/baz", "github.com/becheran/smock/baz"},
	}
	for id, test := range suite {
		t.Run(fmt.Sprintf("%d", id), func(t *testing.T) {
			assert.Equal(t, gomod.PathToUnix(test.result), test.info.ModImportPath(test.path))
		})
	}
}
