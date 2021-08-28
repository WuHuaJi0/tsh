package main

import (
	"bufio"
	"os"
)

func main() {
	prompt()
	reader := bufio.NewReader(os.Stdin)
	for {
		command, _ := reader.ReadString('\n')
		print(command)
		prompt()
	}
}
