package gomod

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

var rootDir = "."

const modName = "go.mod"

type ModInfo struct {
	Path       string
	ModuleName string
}

func FindMod(startFile string) (info ModInfo, err error) {
	dir := path.Dir(startFile)
	for dir != rootDir {
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
		if dir == "/" {
			dir = rootDir
		} else {
			dir = path.Dir(dir)
		}
	}

	return ModInfo{}, fmt.Errorf("module file %s not found", modName)
}

// ModImportPath returns the go import path for the given file path.
func ModImportPath(modInfo *ModInfo, path string) string {
	modInfoPathUnix := strings.ReplaceAll(modInfo.Path, "\\", "/")
	pathUnix := strings.ReplaceAll(path, "\\", "/")
	subPath := strings.TrimPrefix(pathUnix, modInfoPathUnix)
	return fmt.Sprintf("%s%s", modInfo.ModuleName, subPath)
}
