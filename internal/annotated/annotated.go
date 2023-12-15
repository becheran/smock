package annotated

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path"

	"github.com/becheran/smock/internal/generate"
	"github.com/becheran/smock/internal/logger"
	"github.com/becheran/smock/internal/parse"
	"github.com/becheran/smock/internal/pathhelper"
)

// GenerateMocks for interface at file and line
func GenerateMocks(file string, line int) (mockFile string) {
	importPath := parse.ImportPath(file)
	fset := token.NewFileSet()
	logger.Printf("Parse file '%s'\n", file)
	f, err := parser.ParseFile(fset, file, nil, 0)
	if err != nil {
		log.Fatalf("Failed to parse file '%s'. %s", file, err)
	}

	// TODO mod name differ dir name!
	logger.Printf("Add own package %s to imports", importPath)
	f.Imports = append(f.Imports, &ast.ImportSpec{
		Name: &ast.Ident{Name: f.Name.Name},
		Path: &ast.BasicLit{Value: importPath},
	})

	i, err := parse.ParseInterfaceAtPosition(fset, f, line)
	if err != nil {
		log.Fatalf("Failed to parse interface. %s", err)
	}

	mockFilePathCh := make(chan string)
	go func() {
		mockFilePath := pathhelper.MockFilePath(file, i.PackageName, i.Name)
		logger.Printf("Create mock file: '%s'", mockFilePath)
		if err := os.MkdirAll(path.Dir(mockFilePath), os.ModePerm); err != nil {
			log.Fatalf("Failed to create directory '%s'. %s", path.Dir(mockFilePath), err)
		}
		mockFilePathCh <- mockFilePath
	}()
	m, err := generate.GenerateMock(i)
	if err != nil {
		log.Fatalf("Failed to generate mock. %s", err)
	}

	mockFilePath := <-mockFilePathCh
	if err := os.WriteFile(mockFilePath, m, 0644); err != nil {
		log.Fatalf("Failed to write mock file '%s'. %s", mockFilePath, err)
	}
	return mockFilePath
}
