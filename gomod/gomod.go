package gomod

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/becheran/smock/model"
)

var rootDir = "."

const modName = "go.mod"

type ModInfo struct {
	Path       string
	ModuleName string
}

func FindMod(startFile string) (info ModInfo, err error) {
	dir := path.Dir(PathToUnix(startFile))
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
	subPath := strings.TrimPrefix(PathToUnix(dir), PathToUnix(modInfo.Path))
	return fmt.Sprintf("%s%s", modInfo.ModuleName, subPath)
}

// MockDir returns the file used to for the generated mocks.
func (modInfo ModInfo) MockFilePath(filePath, interfaceName string) string {
	modInfoPathUnix := PathToUnix(modInfo.Path)
	subPath := strings.TrimPrefix(PathToUnix(filePath), modInfoPathUnix)
	goIdx := strings.LastIndex(subPath, ".go")
	subPath = fmt.Sprintf("%s_%s%s%s", subPath[:goIdx], strings.ToLower(interfaceName), model.MockPackageSuffix, subPath[goIdx:])
	return fmt.Sprintf("%s/%s%s", modInfoPathUnix, model.MockDir, subPath)
}

func PathToUnix(p string) string {
	return strings.ReplaceAll(p, "\\", "/")
}
