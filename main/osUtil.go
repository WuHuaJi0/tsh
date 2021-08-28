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
