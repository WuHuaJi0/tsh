package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func getPathDirectories() []string {
	getenv := os.Getenv("PATH")
	return strings.Split(getenv, ":")
}

func commandInPath(command string) bool {
	directories := getPathDirectories()
	for i := range directories {
		dir, _ := ioutil.ReadDir(directories[i])
		for j := range dir {
			if dir[j].Name() == command {
				return true
			}
		}
	}
	return false
}

func replaceHomeToTilde(pwd string) string {
	if strings.Contains(pwd, os.Getenv("HOME")) {
		return strings.Replace(pwd, os.Getenv("HOME"), "~", 1)
	}
	return pwd
}

func replaceTildeToHome(pwd string) string {
	if strings.HasPrefix(pwd, "~") {
		return strings.Replace(pwd, "~", os.Getenv("HOME"), 1)
	}
	return pwd
}
