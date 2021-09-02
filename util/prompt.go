package util

import (
	"fmt"
	"os"
	"os/user"
)

/**
 * 用于输出命令行的 Prompt
 */
func Prompt() {
	current, _ := user.Current()
	hostname, _ := os.Hostname()
	pwdOrigin, _ := os.Getwd()
	pwd := replaceHomeToTilde(pwdOrigin)
	fmt.Printf("\033[0;32;34m%s@\033[32m%s\033[m:%s:", current.Username, hostname, pwd)
}
