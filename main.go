package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/KM911/obsidian-execute-code-c-support/util"
)

var (
	cpp_temp = filepath.Join(util.EXECUTE_PATH, "temp.cpp")
	exe_temp = filepath.Join(util.EXECUTE_PATH, "temp.exe")
	CC       = ""
)

func init() {
	// TODO Read config file

	if util.ExecuteSilent("g++ --version") == 0 {
		CC = "g++ "
		return
	}

	if util.ExecuteSilent("clang++ --version") == 0 {
		CC = "clang++ "
		return
	}
	println("No C++ compiler found")
	os.Exit(1)
}

func HelpMessage() {

}
func DisplayArgs() {
	println("args: ", len(os.Args))
	for index, arg := range os.Args {
		println(index, arg)
	}
}

func Execute() {
	// TODO Read compiler flag and could choose compiler

	util.CreatFile(cpp_temp, strings.Join(os.Args[5:], "\n"))
	command := CC + cpp_temp + " -o " + exe_temp + " && " + exe_temp
	util.Execute(command)
	os.Remove(cpp_temp)
	os.Remove(exe_temp)
}
func main() {
	// DisplayArgs()
	argc := len(os.Args)
	switch argc {
	case 1:
		println("no args")
		HelpMessage()
	default:
		Execute()
	}
}

//  Hi guys , I try to write a adaptor for obsidian-execute-code to support C/C++. you could use g++ or clang++ to compile and run your code instead of cling.

// I hope this could be helpful for you.
