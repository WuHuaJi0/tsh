package builtin

import (
	"fmt"
	"os"
	"tinyshell/util"
)

func isBuiltIn(command string) (bool, int) {
	builtInCommands := [...]string{
		"cd",
		"history",
	}
	for i := range builtInCommands {
		if builtInCommands[i] == command {
			return true, i
		}
	}
	return false, -1
}

func Cd(path []string) {
	if len(path) > 1 {
		util.Err("cd failed: more than one args")
	} else if len(path) == 0 || path[0] == "" { //if don't have a path , we jump to home directory.
		os.Chdir(os.Getenv("HOME"))
	} else {
		err := os.Chdir(util.ReplaceTildeToHome(path[0]))
		if err != nil {
			fmt.Println("cd failed:" + "command: no such file or directory")
		}
	}
}
