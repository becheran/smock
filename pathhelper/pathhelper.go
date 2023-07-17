package pathhelper

import (
	"fmt"
	"path"
	"strings"
)

func PathToUnix(p string) string {
	return strings.ReplaceAll(p, "\\", "/")
}

func MockFilePath(origFile, packageName, interfaceName string) (mockFile string) {
	dir := path.Join(path.Dir(path.Dir(origFile)), fmt.Sprintf("%s_mock", packageName))
	file := fmt.Sprintf("%s_%s.go", path.Base(origFile[:len(origFile)-3]), strings.ToLower(interfaceName))

	return path.Join(dir, file)
}
