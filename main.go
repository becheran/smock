// Smock is a mocking library to generate mocks for go interfaces.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strconv"

	"github.com/becheran/smock/internal/annotated"
	"github.com/becheran/smock/internal/logger"
)

func main() {
	dbg := false
	printVersion := false
	flag.BoolVar(&dbg, "debug", false, "print debug information")
	flag.BoolVar(&printVersion, "v", false, "print smock version")
	flag.Parse()

	version := mainVersion()

	if printVersion {
		fmt.Println(version)
		return
	}

	if dbg {
		logger.EnableLogger()
		logger.Printf("Debug mode enabled\n")
	}

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

	annotated.GenerateMocks(file, line, version)
}

func mainVersion() string {
	version := "unknown"
	if info, found := debug.ReadBuildInfo(); found {
		version = info.Main.Version
	}
	return version
}
