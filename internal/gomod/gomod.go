package gomod

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/becheran/smock/internal/logger"
	"github.com/becheran/smock/internal/pathhelper"
)

var rootDir = "."

const MOD_NAME = "go.mod"

type ModInfo struct {
	Path       string
	ModuleName string
}

func FindMod(startPath string) (info ModInfo, err error) {
	logger.Printf("Search for module in '%s'", startPath)
	fileInfo, err := os.Stat(startPath)
	if err != nil {
		return ModInfo{}, err
	}
	dir := pathhelper.PathToUnix(startPath)
	if !fileInfo.IsDir() {
		dir = path.Dir(dir)
	}
	for {
		modFile := path.Join(dir, MOD_NAME)
		if _, existsErr := os.Stat(modFile); existsErr == nil {
			return OpenModFile(modFile)
		}
		if dir == "/" || dir == "" || dir == "." || dir == rootDir {
			return ModInfo{}, fmt.Errorf("module file '%s' for path '%s' not found", MOD_NAME, startPath)
		}
		dir = path.Dir(dir)
	}
}

func OpenModFile(modFile string) (ModInfo, error) {
	logger.Printf("Found mod file '%s'", modFile)
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
		Path:       path.Dir(modFile),
		ModuleName: moduleName,
	}, nil
}

// ModImportPath returns the go import path for the given file path.
func (modInfo ModInfo) ModImportPath(dir string) string {
	subPath := strings.TrimPrefix(pathhelper.PathToUnix(dir), pathhelper.PathToUnix(modInfo.Path))
	return fmt.Sprintf("%s%s", modInfo.ModuleName, subPath)
}
