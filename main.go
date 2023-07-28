package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strconv"

	"github.com/becheran/smock/internal/logger"
	"github.com/becheran/smock/internal/smock"
)

func main() {
	dbg := false
	version := false

	flag.BoolVar(&dbg, "debug", false, "print debug information")
	flag.BoolVar(&version, "v", false, "print smock version")
	flag.Parse()

	if version {
		v := "unknown"
		if info, found := debug.ReadBuildInfo(); found {
			v = info.Main.Version
		}
		fmt.Println(v)
		return
	}

	if dbg {
		logger.SetLogger(log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile))
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

	smock.GenerateMocks(file, line)
}
