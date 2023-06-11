package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path"
	"strconv"

	"github.com/becheran/smock/generate"
	"github.com/becheran/smock/gomod"
	"github.com/becheran/smock/model"
	"github.com/becheran/smock/parse"
)

func main() {
	if os.Getenv("GOLINE") == "" {
		log.Fatal("GOLINE environment variable must be set")
	}
	line, err := strconv.Atoi(os.Getenv("GOLINE"))
	if err != nil {
		log.Fatalf("Failed to parse line number %s. %s", os.Getenv("GOLINE"), err)
	}
	fileName := os.Getenv("GOFILE")
	if fileName == "" {
		log.Fatal("GOFILE environment variable must be set")
	}
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory. %s", err)
	}
	file := fmt.Sprintf("%s/%s", wd, fileName)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, file, nil, 0)
	if err != nil {
		log.Fatalf("Failed to parse file '%s'. %s", fileName, err)
	}

	i, err := parse.ParseInterface(fset, f, line)
	if err != nil {
		log.Fatalf("Failed to parse interface. %s", err)
	}

	modInfo, err := gomod.FindMod(file)
	if err != nil {
		log.Fatalf("Failed to find module. %s", err)
	}
	importPath := modInfo.ModImportPath(path.Dir(file))

	// Add own package as import
	i.Imports = append(i.Imports, model.Import{Path: importPath})

	m, err := generate.GenerateMock(i)
	if err != nil {
		log.Fatalf("Failed to generate mock. %s", err)
	}

	mockFilePath := modInfo.MockFilePath(file)
	if err := os.MkdirAll(path.Dir(mockFilePath), os.ModePerm); err != nil {
		log.Fatalf("Failed to create directory '%s'. %s", path.Dir(mockFilePath), err)
	}
	if err := os.WriteFile(mockFilePath, []byte(m), 0644); err != nil {
		log.Fatalf("Failed to write mock file '%s'. %s", mockFilePath, err)
	}
}
