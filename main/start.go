package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	prompt()
	reader := bufio.NewReader(os.Stdin)
	for {
		//todo: why not syscall.forkexec is not work correctly here.
		line, _ := reader.ReadString('\n')
		command, args := lineToCommand(line)

		if command == "cd" {
			cd(args)
			prompt()
			continue
		}

		_, err := exec.LookPath(command)
		if err != nil {
			fmt.Println("tsh:" + "command not found:" + command)
			prompt()
			continue
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
		prompt()
	}
}
