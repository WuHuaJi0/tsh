package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"tinyshell/builtin"
	"tinyshell/util"
)

func run(cmd util.Command) {

	builtin.Record(cmd.Cmd)

	if cmd.Cmd == "cd" {
		builtin.Cd(cmd.Args)
		return
	}

	if cmd.Cmd == "history" {
		builtin.History()
		return
	}

	if result, index := builtin.IsSearchHistory(cmd.Cmd); result {
		commandString := builtin.GetHistory(index)
		historyCmd, success := util.ParseCommand(commandString)
		//command, args := util.LineToCommand(commandString)
		if !success {
			fmt.Println("tsh: parse command err")
		}
		run(historyCmd)
		return
	}

	_, err := exec.LookPath(cmd.Cmd)
	if err != nil {
		fmt.Println("tsh:" + "command not found:" + cmd.Cmd)
		return
	}

	execCmd := exec.Command(cmd.Cmd, cmd.Args...)

	util.Redirection(execCmd, cmd)
	execCmd.Run()

	if err != nil {
		code := err.(*exec.ExitError).ProcessState.ExitCode()
		// Somethings many command may return 1 when not have enough args or received error args
		// for example when just typing "git"  or  "git blabla" in shell.
		if code != 1 {
			fmt.Println("Execute Command failed:" + err.Error())
		}
	}

}

func main() {
	util.Prompt()
	reader := bufio.NewReader(os.Stdin)
	for {
		//todo: why not syscall.forkexec is not work correctly here.
		line, _ := reader.ReadString('\n')
		//command, args := util.LineToCommand(line)

		cmd, success := util.ParseCommand(line)
		if !success {
			fmt.Println("tsh: parse command err")
			util.Prompt()
			continue
		}

		if cmd.Cmd == "exit" {
			println("Bye...")
			os.Exit(0)
		}

		run(cmd)

		util.Prompt()
	}
}
