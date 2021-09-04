package util

import (
	"fmt"
	"strings"
)

type Command struct {
	Cmd          string
	Args         []string
	Stdin        string
	StdinAppend  bool
	Stdout       string
	StdoutAppend bool
	Stderr       string
	StderrAppend bool
}

func LineToCommand(line string) (string, []string) {
	line = strings.Replace(strings.TrimSpace(line), "\n", "", -1) // remove the last \n
	command := strings.Split(line, " ")
	if len(command) > 1 {
		return command[0], command[1:] //if this command has more than one args.
	}
	return command[0], nil
}

func ParseCommand(line string) (Command, bool) {
	var command Command
	command = Command{}
	args := []string{}

	line = strings.Replace(strings.TrimSpace(line), "\n", "", -1) // remove the last \n
	strArr := strings.Split(line, " ")
	command.Cmd = strArr[0]

	if len(strArr) > 1 {
		for i := 1; i < len(strArr); i++ {
			redirectType, isAppend := RedirectionType(strArr[i])
			if (redirectType != "") && len(strArr) <= i {
				if len(strArr) <= i {
					fmt.Println("tsh: redirection doesn't have the dist file")
					return command, false
				}
			} else if redirectType == "stdin" {
				command.Stdin = strArr[i+1]
				command.StdinAppend = isAppend
				i++
			} else if redirectType == "stdout" {
				command.Stdout = strArr[i+1]
				command.StdoutAppend = isAppend
				i++
			} else {
				args = append(args, strArr[i])
			}
		}
	}
	command.Args = args

	return command, true
}
