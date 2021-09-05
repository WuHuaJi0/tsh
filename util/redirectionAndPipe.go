/**
 * 处理管道和重定向的一些函数
 */
package util

import (
	"bufio"
	"os"
	"os/exec"
)

func RedirectionType(args string) (string, bool) {
	if args == "<" {
		return "stdin", false
	} else if args == "<<" {
		return "stdin", true
	} else if args == ">" {
		return "stdout", false
	} else if args == ">>" {
		return "stdout", true
	}
	return "", false
}

func Redirection(execCmd *exec.Cmd, cmd Command) {
	//处理标准输入
	if cmd.Stdin != "" {
		flagStdin := os.O_RDONLY
		if cmd.StdinAppend {
			flagStdin = flagStdin | os.O_APPEND
		}
		file, err := os.OpenFile(cmd.Stdin, flagStdin, 0777)
		if err != nil {
			Err("tsh: read stdin failed!")
			return
		}
		execCmd.Stdin = file
	} else {
		execCmd.Stdin = os.Stdin
	}

	//处理标准输出
	if cmd.Stdout != "" {
		flagStdout := os.O_WRONLY | os.O_CREATE
		if cmd.StdoutAppend {
			flagStdout = flagStdout | os.O_APPEND
		}
		file, err := os.OpenFile(cmd.Stdout, flagStdout, 0666)
		if err != nil {
			Err("open stdout file failed!")
			return
		}
		execCmd.Stdout = file
	} else {
		execCmd.Stdout = os.Stdout
	}

	//todo: 需要处理标准错误
	if cmd.Stderr != "" {

	} else {
		execCmd.Stderr = os.Stderr
	}
}

/**
 * 运行当前命令，并且把当前命令的输出作为下个命令的输入
 */
func RunAndSetPipe(current *exec.Cmd, next *exec.Cmd) {
	current.Stdout = nil
	next.Stdin = nil
	pipe, _ := current.StdoutPipe()
	current.Start()
	outputBuf1 := bufio.NewReader(pipe)
	stdinPipe, _ := next.StdinPipe()
	outputBuf1.WriteTo(stdinPipe)
	current.Wait()
	stdinPipe.Close()
}
