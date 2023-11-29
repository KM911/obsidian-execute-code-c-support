package util

import "path/filepath"

var (
	EXECUTE_PATH = ExecutePath()
	CMD_PATH     = CmdPath()
)

func init() {
	FileLogger(filepath.Join(EXECUTE_PATH, "log.txt"))

}
