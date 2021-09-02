package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"tinyshell/builtin"
	"tinyshell/util"
)

func run(command string, args []string) {

	builtin.Record(command)

	if command == "cd" {
		builtin.Cd(args)
		return
	}

	if command == "history" {
		builtin.History()
		return
	}

	if result, index := builtin.IsSearchHistory(command); result {
		commandString := builtin.GetHistory(index)
		command, args := util.LineToCommand(commandString)
		run(command, args)
		return
	}

	_, err := exec.LookPath(command)
	if err != nil {
		fmt.Println("tsh:" + "command not found:" + command)
		return
	}

	cmd := exec.Command(command, args...)

	output, err := cmd.Output()
	if err != nil {
		code := err.(*exec.ExitError).ProcessState.ExitCode()
		// Somethings many command may return 1 when not have enough args or received error args
		// for example when just typing "git"  or  "git blabla" in shell.
		if code != 1 {
			fmt.Println("Execute Command failed:" + err.Error())
		}
	}
	fmt.Print(string(output))
}

func main() {
	util.Prompt()
	reader := bufio.NewReader(os.Stdin)
	for {
		//todo: why not syscall.forkexec is not work correctly here.
		line, _ := reader.ReadString('\n')
		command, args := util.LineToCommand(line)

		run(command, args)

		util.Prompt()
	}
}
