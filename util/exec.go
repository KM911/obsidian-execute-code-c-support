package util

import (
	"os"
	"os/exec"
	"runtime"
)

var (
	CreateCommand func(command string) *exec.Cmd
)

func init() {
	if runtime.GOOS == "windows" {
		CreateCommand = createCmd
	} else {
		CreateCommand = createBash
	}
}

func redriectStd(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
}

func createCmd(command string) *exec.Cmd {
	return exec.Command("cmd", "/C", command)
}

func createBash(command string) *exec.Cmd {
	return exec.Command("bash", "-c", command)
}

func ExecuteCommand(command string) int {
	cmdExecutor := CreateCommand(command)
	redriectStd(cmdExecutor)
	cmdExecutor.Run()
	return cmdExecutor.ProcessState.ExitCode()
}

func ExecuteCommandSilent(command string) int {
	cmdExecutor := CreateCommand(command)
	cmdExecutor.Run()
	return cmdExecutor.ProcessState.ExitCode()
}

func ExecuteCommandResult(command string) string {
	cmdExecutor := CreateCommand(command)
	redriectStd(cmdExecutor)
	output, _ := cmdExecutor.Output()
	return string(output)
}

func ExecuteCommandSilentResult(command string) string {
	cmdExecutor := CreateCommand(command)
	output, _ := cmdExecutor.Output()
	return string(output)
}

func KillByName(name string) {
	ExecuteCommandSilent("taskkill /F /IM " + name)
}
