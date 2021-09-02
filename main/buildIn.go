package main

import (
	"fmt"
	"os"
)

func isBuiltIn(command string) (bool, int) {
	builtInCommands := [...]string{
		"cd",
	}
	for i := range builtInCommands {
		if builtInCommands[i] == command {
			return true, i
		}
	}
	return false, -1
}

func cd(path []string) {
	if len(path) > 1 {
		fmt.Println("cd failed: more than one args")
	} else if len(path) == 0 || path[0] == "" { //if don't have a path , we jump to home directory.
		os.Chdir(os.Getenv("HOME"))
	} else {
		err := os.Chdir(replaceTildeToHome(path[0]))
		if err != nil {
			fmt.Println("cd failed:" + "command: no such file or directory")
		}
	}
}
