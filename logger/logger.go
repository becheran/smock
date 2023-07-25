package logger

import (
	"fmt"
	"log"
)

var logger *log.Logger = nil

func SetLogger(l *log.Logger) {
	logger = l
}

func Printf(format string, v ...interface{}) {
	if logger == nil {
		return
	}
	if err := logger.Output(2, fmt.Sprintf(format, v...)); err != nil {
		panic(err)
	}
}
