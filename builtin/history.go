package builtin

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

// Record the command to a history file
func Record(command string) {
	path := os.Getenv("HOME") + "/.tsh_history"
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil && os.IsNotExist(err) {
		_, createErr := os.Create(path)
		if createErr != nil {
			fmt.Println("tsh: create .tsh_history failed！")
			return
		}
	}

	timestamp := time.Now().Unix()
	s := strconv.FormatInt(timestamp, 10)

	_, err = file.WriteString(s + ":" + command + "\n")

	if err != nil {
		fmt.Println("tsh: record history failed!")
	}
}

// read all History in History file
func History() {
	path := os.Getenv("HOME") + "/.tsh_history"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("tsh: can not read .tsh_history")
		return
	}
	all, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("tsh: can not read .tsh_history")
		return
	}
	commands := strings.Split(strings.TrimRight(string(all[:]), "\n"), "\n")
	for i := range commands {
		fmt.Println(strconv.Itoa(i+1) + " " + strings.Split(commands[i], ":")[1])
	}
}

//checkout if a command is like !number, eg: !123
func IsSearchHistory(command string) (bool, int) {
	if !strings.HasPrefix(command, "!") {
		return false, -1
	}
	number := strings.Replace(command, "!", "", 1)
	if _, err := strconv.Atoi(number); err == nil {
		index, _ := strconv.Atoi(number)
		return true, index
	}
	return false, 0
}

func GetHistory(index int) string {
	file, _ := os.Open(os.Getenv("HOME") + "/.tsh_history")
	scanner := bufio.NewScanner(file)
	current := 1
	for scanner.Scan() {
		if current == index {
			text := scanner.Text()
			split := strings.Split(text, ":")
			return split[1]
		}
		current++
	}
	return "tsh: can not find such history"
}
