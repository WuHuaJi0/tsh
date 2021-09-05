package util

import (
	"fmt"
	"os/exec"
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

/**
 * 把通过管道符连接的命令拆分
 */
func ParsePipe(line string) []string {
	line = strings.Replace(strings.TrimSpace(line), "\n", "", -1) // remove the last \n
	strArr := strings.Split(line, "|")
	return strArr
}

func ParseCommand(line string) (Command, bool) {
	var command Command
	command = Command{}
	args := []string{}

	strArr := strings.Split(strings.TrimSpace(line), " ")
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

/**
 * 把一行字符串，拆分成多个命令
 */
func LineToCommand(line string) ([]Command, bool) {
	cmdList := []Command{}
	pipe := ParsePipe(line)
	for i := range pipe {
		cmd, success := ParseCommand(pipe[i])
		if !success {
			Err("parse command err")
			return cmdList, false
		}
		cmdList = append(cmdList, cmd)
	}
	return cmdList, true
}

/**
 * 把 command 结构体数组，
 * 转换为 exec.Command 数组
 */
func CommandsToExecCommands(cmdList []Command) ([]*exec.Cmd, bool) {
	execCmdList := []*exec.Cmd{}
	for _, cmd := range cmdList {
		_, err := exec.LookPath(cmd.Cmd)
		if err != nil {
			Err("command not found:" + cmd.Cmd)
			return nil, false
		}
		execCmd := exec.Command(cmd.Cmd, cmd.Args...)
		Redirection(execCmd, cmd)
		execCmdList = append(execCmdList, execCmd)
	}
	return execCmdList, true
}
