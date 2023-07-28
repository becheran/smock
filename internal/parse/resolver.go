package parse

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"io/fs"
	"strings"

	"github.com/becheran/smock/internal/logger"
	"github.com/becheran/smock/internal/model"
)

func parsePackage(packageId string, imports []*ast.ImportSpec, dir string) (pkg *ast.Package, err error) {
	logger.Printf("Resolve package %s in %s", packageId, dir)

	for _, astImp := range imports {
		imp := model.ImportFromAst(astImp)
		if imp.ImportName() == packageId {
			buildPkg, err := build.Import(imp.Path, "./", build.FindOnly)
			if err != nil {
				return nil, err
			}
			dir = buildPkg.Dir
			break
		}
	}

	noTestFileFilter := func(fi fs.FileInfo) bool { return !strings.HasSuffix(fi.Name(), "_test.go") }
	parseRes, err := parser.ParseDir(token.NewFileSet(), dir, noTestFileFilter, parser.Mode(0))
	if err != nil {
		return nil, err
	}

	if len(parseRes) != 1 {
		return nil, fmt.Errorf("expected one package, but got %d", len(parseRes))
	}

	for _, p := range parseRes {
		pkg = p
	}

	return pkg, nil
}
