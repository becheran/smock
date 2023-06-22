package logger

import (
	"fmt"
	"io"
	"log"
)

var Logger *log.Logger = log.New(io.Discard, "", 0)

func Printf(format string, v ...interface{}) {
	if err := Logger.Output(2, fmt.Sprintf(format, v...)); err != nil {
		panic(err)
	}
}
