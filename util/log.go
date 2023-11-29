package util

import (
	"log"
	"os"
)

func FileLogger(src string) {
	logFile, err := os.OpenFile(src, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
}

func FileLoggerClose() {
	log.SetOutput(os.Stdout)
}

func FileLoggerPrintln(v ...interface{}) {
	log.Println(v...)
}
