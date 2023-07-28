package gomod

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/becheran/smock/internal/pathhelper"
)

var rootDir = "."

const modName = "go.mod"

type ModInfo struct {
	Path       string
	ModuleName string
}

func FindMod(startFile string) (info ModInfo, err error) {
	dir := path.Dir(pathhelper.PathToUnix(startFile))
	for {
		modFile := path.Join(dir, modName)
		if _, existsErr := os.Stat(modFile); existsErr == nil {
			file, err := os.Open(modFile)
			if err != nil {
				return ModInfo{}, err
			}
			scanner := bufio.NewScanner(file)
			scanner.Scan()
			firstLine := scanner.Text()
			file.Close()

			moduleName := strings.TrimPrefix(firstLine, "module ")
			if moduleName == "" {
				return ModInfo{}, fmt.Errorf("failed to parse module name in %s", modFile)
			}
			return ModInfo{
				Path:       dir,
				ModuleName: moduleName,
			}, nil
		}
		if dir == "/" || dir == "" || dir == "." || dir == rootDir {
			return ModInfo{}, fmt.Errorf("module file '%s' for path '%s' not found", modName, startFile)
		}
		dir = path.Dir(dir)
	}
}

// ModImportPath returns the go import path for the given file path.
func (modInfo ModInfo) ModImportPath(dir string) string {
	subPath := strings.TrimPrefix(pathhelper.PathToUnix(dir), pathhelper.PathToUnix(modInfo.Path))
	return fmt.Sprintf("%s%s", modInfo.ModuleName, subPath)
}
