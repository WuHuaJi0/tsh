package util

import "strings"

func LineToCommand(line string) (string, []string) {
	line = strings.Replace(line, "\n", "", -1) // remove the last \n
	command := strings.Split(line, " ")
	if len(command) > 1 {
		return command[0], command[1:] //if this command has more than one args.
	}
	return command[0], nil
}
