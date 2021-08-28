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
		cmd := exec.Command(command, args...)
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("Execute Command failed:" + err.Error())
			return
		}
		fmt.Println(string(output))
		prompt()
	}
}
