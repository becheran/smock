// Package smock exposes methods to generate new mocks for interfaces in a go project.
package smock

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/becheran/smock/internal/generate"
	"github.com/becheran/smock/internal/gomod"
	"github.com/becheran/smock/internal/logger"
	"github.com/becheran/smock/internal/parse"
	"github.com/becheran/smock/internal/pathhelper"
	"github.com/becheran/wildmatch-go"
)

type opt struct {
	Unexported    bool
	Debug         bool
	InterfaceList []*wildmatch.WildMatch
	IsAllowList   bool
}

func skip(name string, list []*wildmatch.WildMatch, isAllowList, allowUnexported bool) (skip bool) {
	if !token.IsExported(name) && !allowUnexported {
		return true
	}
	skip = isAllowList
	for _, m := range list {
		if m.IsMatch(name) {
			if isAllowList {
				return false
			} else {
				return true
			}
		}
	}
	return skip
}

// Option which can be passed to [GenerateMocks] using one of the constructors.
type Option func(*opt)

// WithDebugLog enables debug logging which prints debug logs to the console
func WithDebugLog() Option {
	return func(o *opt) {
		o.Debug = true
	}
}

// WithInterfaceNameAllowList sets a list of interface names which shall be used for mock creation.
// Is incompatible with the [WithInterfaceNameDenyList] option.
//
// Uses the wildcard syntax described in http://github.com/becheran/wildmatch-go for string matches.
func WithInterfaceNameAllowList(allow ...string) Option {
	return func(o *opt) {
		for _, a := range allow {
			o.InterfaceList = append(o.InterfaceList, wildmatch.NewWildMatch(a))
		}
		o.IsAllowList = true
	}
}

// WithInterfaceNameDenyList sets a list of interface names which shall not be used for mock creation.
// Is incompatible with the [WithInterfaceNameAllowList] option.
//
// Uses the wildcard syntax described in http://github.com/becheran/wildmatch-go for string matches.
func WithInterfaceNameDenyList(deny ...string) Option {
	return func(o *opt) {
		for _, d := range deny {
			o.InterfaceList = append(o.InterfaceList, wildmatch.NewWildMatch(d))
		}
		o.IsAllowList = false
	}
}

// WithUnexportedInterfaces enables generation of unexported interfaces
func WithUnexportedInterfaces() Option {
	return func(o *opt) {
		o.Unexported = true
	}
}

// GenerateMocks creates mock objects for all interfaces found in the module.
// Will search for a "go.mod" file traversing up the filesystem. Fails if no mod file is found.
//
// Returns a list of file paths to the generated mock objects.
// Optionally pass multiple [Option] as parameter.
//
// Is meant to be used within a short go main function which will execute [GenerateMocks].
// Annotate with '//go:generate go run ./' to allow mock generation via the 'go generate' command.
// The advantage of doing so is that no additional software needs to be installed to generate the mocks
// since smock is already included in the module as library. This simplifies the project setup because no
// additional binaries needs to be installed on the machine which generates the mocks. The mock files might then also
// be excluded from your repository by a adding '*/**/*_mock' to the '.gitignore' file.
//
// The following code snippet shows how smock can be used as a small binary within a go project:
//
//	package main
//
//	import (
//	    "github.com/becheran/smock/smock"
//	)
//
//	//go:generate go run ./
//	func main() {
//	    smock.GenerateMocks()
//	}
func GenerateMocks(options ...Option) (mockFilePaths []string) {
	opt := &opt{}
	for _, o := range options {
		o(opt)
	}
	if opt.Debug {
		logger.EnableLogger()
	}

	fset := token.NewFileSet()

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	modInfo, err := gomod.FindMod(wd)
	if err != nil {
		log.Fatalf("Failed to open module. %s", err)
	}

	err = filepath.Walk(modInfo.Path,
		func(filePath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				return nil
			}
			if strings.HasSuffix(info.Name(), "_test") || strings.HasSuffix(info.Name(), "_mock") || info.Name() == ".git" {
				return filepath.SkipDir
			}

			logger.Printf("Parse dir '%s'", filePath)
			dir, err := parser.ParseDir(fset, filePath, nil, 0)
			if err != nil {
				log.Fatalf("Failed to parse dir '%s'. %s", filePath, err)
			}
			for _, pkg := range dir {
				for fileName, file := range pkg.Files {
					for _, d := range file.Decls {
						ts, err := parse.GetTypeSpec(d)
						if err != nil {
							continue
						}

						importPath := modInfo.ModImportPath(path.Dir(pathhelper.PathToUnix(fileName)))
						// TODO mod name differ dir name!
						logger.Printf("Add own package %s to imports", importPath)
						file.Imports = append(file.Imports, &ast.ImportSpec{
							Name: &ast.Ident{Name: pkg.Name},
							Path: &ast.BasicLit{Value: importPath},
						})

						i, err := parse.ParseInterface(ts, pkg.Name, fileName, file.Imports)
						if err != nil {
							continue
						}

						logger.Println("Found interface", i.Name)
						if skip(i.Name, opt.InterfaceList, opt.IsAllowList, opt.Unexported) {
							logger.Println("Skip interface", i.Name)
							continue
						}

						m, err := generate.GenerateMock(i)
						if err != nil {
							log.Fatalf("Failed to generate mock. %s", err)
						}

						mockFilePath := pathhelper.MockFilePath(fileName, i.PackageName, i.Name)
						logger.Printf("Create mock file: '%s'", mockFilePath)
						if err := os.MkdirAll(path.Dir(mockFilePath), os.ModePerm); err != nil {
							log.Fatalf("Failed to create directory '%s'. %s", path.Dir(mockFilePath), err)
						}
						if err := os.WriteFile(mockFilePath, m, 0644); err != nil {
							log.Fatalf("Failed to write mock file '%s'. %s", mockFilePath, err)
						}
						mockFilePaths = append(mockFilePaths, mockFilePath)
					}
				}
			}
			return nil
		})
	if err != nil {
		panic(err)
	}
	return mockFilePaths
}
