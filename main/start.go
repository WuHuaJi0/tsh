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

		if !commandInPath(command) {
			fmt.Println("tsh:" + "command not found:" + command)
			prompt()
			continue
		}
		cmd := exec.Command(command, args...)
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("Execute Command failed:" + err.Error())
			prompt()
			continue
		}
		fmt.Println(string(output))
		prompt()
	}
}
