package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/KM911/obsidian-execute-code-c-support/util"
)

var (
	CppTemp       = filepath.Join(util.ExecuteDirectory, "temp.cpp")
	ExeTemp       = filepath.Join(util.ExecuteDirectory, "temp.exe")
	StringBuilder = strings.Builder{}
	CC            = ""
)

func init() {

	// detect c compiler
	if util.ExecuteCommandSilent("clang++ --version") == 0 {
		CC = "clang++ "
		return
	}

	if util.ExecuteCommandSilent("g++ --version") == 0 {
		CC = "g++ "
		return
	}

	println("No C++ compiler found")
	os.Exit(1)
}

func helpMessage() {

}
func displayArgs() {
	for index, arg := range os.Args {
		println(index, arg)
	}
}

func Execute() {
	// TODO Read compiler flag and could choose compiler

	index := 2
	for !strings.HasPrefix(os.Args[index], "#") {
		index++
	}
	util.CreatFile(CppTemp, strings.Join(os.Args[index:], "\n"))
	command := CC + CppTemp + " -o " + ExeTemp + " && " + ExeTemp
	util.ExecuteCommand(command)
	os.Remove(CppTemp)
	os.Remove(ExeTemp)
}

func main() {
	argc := len(os.Args)
	switch argc {
	case 1:
		println("no args")
		helpMessage()
	default:
		Execute()
	}
}
