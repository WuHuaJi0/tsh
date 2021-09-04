package util

import (
	"os"
	"strings"
)

func replaceHomeToTilde(pwd string) string {
	if strings.Contains(pwd, os.Getenv("HOME")) {
		return strings.Replace(pwd, os.Getenv("HOME"), "~", 1)
	}
	return pwd
}

func ReplaceTildeToHome(pwd string) string {
	if strings.HasPrefix(pwd, "~") {
		return strings.Replace(pwd, "~", os.Getenv("HOME"), 1)
	}
	return pwd
}
