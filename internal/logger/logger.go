package logger

import (
	"fmt"
	"log"
	"os"
)

var logger *log.Logger = nil

func EnableLogger() {
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func Printf(format string, v ...any) {
	if logger == nil {
		return
	}
	if err := logger.Output(2, fmt.Sprintf(format, v...)); err != nil {
		panic(err)
	}
}

func Println(format ...any) {
	if logger == nil {
		return
	}
	if err := logger.Output(2, fmt.Sprintln(format...)); err != nil {
		panic(err)
	}
}
