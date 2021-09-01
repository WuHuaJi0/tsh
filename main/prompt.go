package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

func replaceHome(pwd string) string {
	if strings.Contains(pwd, os.Getenv("HOME")) {
		return strings.Replace(pwd, os.Getenv("HOME"), "~", 1)
	}
	return pwd
}

/**
 * 用于输出命令行的 prompt
 */
func prompt() {
	current, _ := user.Current()
	hostname, _ := os.Hostname()
	pwdOrigin, _ := os.Getwd()
	pwd := replaceHome(pwdOrigin)
	fmt.Printf("\033[0;32;34m%s@\033[32m%s\033[m:%s:", current.Username, hostname, pwd)
}
