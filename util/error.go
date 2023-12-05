package util

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
)

func FileLogger(src string) {
	logFile, err := os.OpenFile(src, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
}

type ErrorMessage struct {
	Type int
	Msg  string
}

func EmitError(n int, msg string) {
	panic(ErrorMessage{
		Type: n,
		Msg:  msg,
	})
}

func Recover(errorHandler func(ErrorMessage)) {
	if errMsg := recover(); errMsg != nil {
		errorHandler(errMsg.(ErrorMessage))
		log.Println(errMsg.(ErrorMessage).Msg)
		log.Println(string(debug.Stack()))
	}
}

func ErrorHandler(err ErrorMessage) {
	switch err.Type {
	case 0:
		fmt.Println(err.Msg)

	// usually be -1
	default:
		// println("panic")
		println("\nerror type : ", err.Type, "\n\nerror message : ", err.Msg)
	}
}

var (
	MessageList []ErrorMessage
)
