package smock

import (
	"go/parser"
	"go/token"
	"log"
	"os"
	"path"

	"github.com/becheran/smock/generate"
	"github.com/becheran/smock/gomod"
	"github.com/becheran/smock/logger"
	"github.com/becheran/smock/model"
	"github.com/becheran/smock/parse"
	"github.com/becheran/smock/pathhelper"
)

// GenerateMocks for interface at file and line
func GenerateMocks(file string, line int) (mockFile string) {
	fset := token.NewFileSet()
	logger.Printf("Parse file '%s'\n", file)
	f, err := parser.ParseFile(fset, file, nil, 0)
	if err != nil {
		log.Fatalf("Failed to parse file '%s'. %s", file, err)
	}

	i, err := parse.ParseInterfaceAtPosition(fset, f, line)
	if err != nil {
		log.Fatalf("Failed to parse interface. %s", err)
	}

	modInfo, err := gomod.FindMod(file)
	if err != nil {
		log.Fatalf("Failed to find module. %s", err)
	}
	importPath := modInfo.ModImportPath(path.Dir(pathhelper.PathToUnix(file)))

	logger.Printf("Add own package to imports")
	i.Imports = append(i.Imports, model.Import{Path: importPath})

	m, err := generate.GenerateMock(i)
	if err != nil {
		log.Fatalf("Failed to generate mock. %s", err)
	}

	mockFilePath := pathhelper.MockFilePath(file, i.PackageName, i.Name)
	logger.Printf("Create mock file: '%s'", mockFilePath)
	if err := os.MkdirAll(path.Dir(mockFilePath), os.ModePerm); err != nil {
		log.Fatalf("Failed to create directory '%s'. %s", path.Dir(mockFilePath), err)
	}
	if err := os.WriteFile(mockFilePath, []byte(m), 0644); err != nil {
		log.Fatalf("Failed to write mock file '%s'. %s", mockFilePath, err)
	}
	return mockFilePath
}