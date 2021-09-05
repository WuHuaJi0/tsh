package main

import (
	"bufio"
	"os"
	"tinyshell/util"
)

//func run(cmd util.Command) (*exec.Cmd, bool) {
//builtin.Record(cmd.Cmd)
////如果是要退出
//if cmd.Cmd == "exit" {
//	fmt.Println("Bye...")
//	os.Exit(0)
//}
//
//if cmd.Cmd == "cd" {
//	builtin.Cd(cmd.Args)
//	return
//}
//
//if cmd.Cmd == "history" {
//	builtin.History()
//	return
//}
//if result, index := builtin.IsSearchHistory(cmd.Cmd); result {
//	commandString := builtin.GetHistory(index)
//	historyCmdList, success := util.LineToCommand(commandString)
//	if !success {
//		util.Err("parse command err")
//}
//run(historyCmdList)
//return
//}

//}

func main() {
	util.Prompt()
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _ := reader.ReadString('\n')
		cmdList, success := util.LineToCommand(line)
		if !success {
			util.Err("parse command err")
			util.Prompt()
			continue
		}

		execCmdList, success := util.CommandsToExecCommands(cmdList)
		if !success {
			util.Prompt()
			continue
		}

		count := len(execCmdList)
		for i, current := range execCmdList {
			if count > 1 && i <= count-2 {
				util.RunAndSetPipe(current, execCmdList[i+1])
			} else {
				current.Run()
			}
		}
		util.Prompt()
	}
}
