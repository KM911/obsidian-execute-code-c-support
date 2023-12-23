package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/KM911/obsidian-execute-code-c-support/util"
)

const (
	VERSION = "0.04"
)

var (
	CppTemp       = filepath.Join(util.ExecuteDirectory, "temp.cpp")
	ExeTemp       = filepath.Join(util.ExecuteDirectory, "temp.exe")
	StringBuilder = strings.Builder{}
	CC            = ""
)

func init() {
	util.FileLogger(filepath.Join(util.ExecuteDirectory, "obsidian-execute-code-c-support.log"))

	// detect c compiler

	if util.ExecuteCommandSilent("g++ --version") == 0 {
		CC = "g++ "
		return
	}

	if util.ExecuteCommandSilent("clang++ --version") == 0 {
		CC = "clang++ "
		return
	}

	println("No C++ compiler found")
	os.Exit(1)
}

func helpMessage() {
	fmt.Println("version : ", VERSION)
	fmt.Println("author : KM911")
	fmt.Println("github : https://github.com/KM911/obsidian-execute-code-c-support")
}
func displayArgs() {
	for index, arg := range os.Args {
		println(index, arg)
	}
}

func Execute() {
	// TODO Read compiler flag and could choose compiler

	index := 2
	for index < len(os.Args) && !strings.HasPrefix(os.Args[index], "#") {
		index++
	}
	if index == len(os.Args) {
		fmt.Println("Please disable the option \"Use main function\"")
		log.Fatal("out of args range")
		return
	}
	// to prevent the previous program do not exit
	util.KillByName("temp.exe")
	util.CreatFile(CppTemp, strings.Join(os.Args[index:], "\n"))
	command := CC + CppTemp + " -w -o " + ExeTemp + " && " + ExeTemp
	print("\u200b")
	util.ExecuteCommand(command)
	os.Remove(CppTemp)
	os.Remove(ExeTemp)
}

func main() {
	argc := len(os.Args)
	//displayArgs()
	switch argc {
	case 1:
		helpMessage()
	default:
		Execute()
	}
}
