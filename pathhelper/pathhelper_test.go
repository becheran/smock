package pathhelper_test

import (
	"fmt"
	"testing"

	"github.com/becheran/smock/pathhelper"
	"github.com/stretchr/testify/assert"
)

func TestMockFilePath(t *testing.T) {
	var suite = []struct {
		origFilePath  string
		modName       string
		interfaceName string
		expdMockFile  string
	}{
		{"/foo/bar.go", "foo", "MockMe", "/foo_mock/bar_mockme.go"},
		{"/foo/baz/bar.go", "baz", "MockMe", "/foo/baz_mock/bar_mockme.go"},
		{"/foo/main.go", "other", "MockMe", "/other_mock/main_mockme.go"},
	}
	for id, test := range suite {
		t.Run(fmt.Sprintf("%d", id), func(t *testing.T) {
			assert.Equal(t, test.expdMockFile, pathhelper.MockFilePath(test.origFilePath, test.modName, test.interfaceName))
		})
	}
}
