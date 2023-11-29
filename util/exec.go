package util

import (
	"os"
	"os/exec"
)

func Execute(command string) {
	cmd_executor := exec.Command("cmd", "/C", command)
	cmd_executor.Stdout = os.Stdout
	cmd_executor.Stdin = os.Stdin
	cmd_executor.Stderr = os.Stderr
	cmd_executor.Run()
}
