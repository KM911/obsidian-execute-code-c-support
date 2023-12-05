package util

import (
	"os"
	"path/filepath"
	"strings"
)

var (
	WorkingDirectory = workingDirectory()
	ExecuteDirectory = executorDirectory()
	TempDirectory    = os.TempDir()
	StringBuilder    = strings.Builder{}
)

/*
// e.g.
// a.txt -> /home/xxx/a.txt
// Makefile -> D:/SOFT/CODE/Makefile
*/
func AbsPath(file_ string) string {
	abs, _ := filepath.Abs(file_)
	return abs
}

/*
e.g.
Pwd() -> /home/xxx
Pwd() -> D:/SOFT/CODE
*/
func workingDirectory() string {
	dir, _ := os.Getwd()
	return dir
}

func executorDirectory() string {
	path, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(path)
}

/*
cover windows path to wsl path
D:\soft\code --> /mnt/d/soft/code
*/
func WslPathConvent(path string) string {
	StringBuilder.Reset()
	StringBuilder.WriteString("/mnt/")
	StringBuilder.WriteString(strings.ToLower(string(path[0])))
	StringBuilder.WriteString(path[3:])
	res := StringBuilder.String()
	StringBuilder.Reset()
	return res
}
