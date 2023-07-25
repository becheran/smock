package test_test

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/becheran/smock/logger"
	"github.com/becheran/smock/smock"
	"github.com/stretchr/testify/assert"
)

const testPackagePath = "./test_package"
const testGoldenFileDir = "golden_test"

var generate = flag.Bool("generate", false, "generate golden files")

// BenchmarkGenerate-12    	     178	   6704438 ns/op	 1758730 B/op	   39041 allocs/op
func BenchmarkGenerate(b *testing.B) {
	os.Chdir(testPackagePath)
	os.RemoveAll("./test_package/testpackage_mock")
	interfaces := getAnnotatedInterfaces()

	for i := 0; i < b.N; i++ {
		for _, i := range interfaces {
			smock.GenerateMocks(i.File, i.Line)
		}
	}
}

func TestGenerate(t *testing.T) {
	logger.SetLogger(log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile))
	os.Chdir(testPackagePath)
	os.RemoveAll("./test_package/testpackage_mock")
	for _, i := range getAnnotatedInterfaces() {
		fmt.Printf("Generate mocks for %s:%d\n", i.File, i.Line)
		file := smock.GenerateMocks(i.File, i.Line)
		source, err := os.Open(file)
		if err != nil {
			panic(err)
		}

		goldenFilePath := goldenFilePath(file)
		if *generate {
			fmt.Printf("Generated mocks for %s:%d in %s\n", i.File, i.Line, goldenFilePath)
			destination, err := os.Create(goldenFilePath)
			if err != nil {
				panic(err)
			}
			_, err = io.Copy(destination, source)
			if err != nil {
				panic(err)
			}
			destination.Close()
		} else {
			golden, err := os.ReadFile(goldenFilePath)
			if err != nil {
				t.Fatalf("Failed to read golden file %s. Might need generate first. %s", goldenFilePath, err)
			}
			generated, err := io.ReadAll(source)
			assert.Nil(t, err)
			assert.Equal(t, string(golden), string(generated))
		}
		source.Close()
	}
}

func goldenFilePath(file string) string {
	fileName := path.Base(file)
	goldenDir := path.Dir(path.Dir(file)) + "/" + testGoldenFileDir + "/"
	if err := os.MkdirAll(goldenDir, os.ModePerm); err != nil {
		panic(err)
	}
	return goldenDir + fileName
}

type fileLocation struct {
	File string
	Line int
}

func getAnnotatedInterfaces() (loc []fileLocation) {
	err := filepath.WalkDir("./", func(path string, d fs.DirEntry, err error) error {
		if d.Type().IsRegular() && strings.HasSuffix(d.Name(), ".go") && !strings.HasSuffix(d.Name(), "_test.go") {
			r, err := os.Open(path)
			if err != nil {
				panic(err)
			}
			scanner := bufio.NewScanner(r)
			scanner.Split(bufio.ScanLines)

			line := 0
			for scanner.Scan() {
				line++
				if strings.HasPrefix(scanner.Text(), "//go:generate smock") {
					absPath, err := filepath.Abs(path)
					if err != nil {
						panic(err)
					}
					loc = append(loc, fileLocation{File: absPath, Line: line})
				}
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return loc
}
