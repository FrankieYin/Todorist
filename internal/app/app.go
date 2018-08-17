package app

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/mitchellh/go-homedir"
	"github.com/FrankieYin/todo/internal/util"
	"github.com/FrankieYin/todo/internal/data"
)

var home string
var todoDir string
var todoJsonFilename string
var archJsonFilename string
var projJsonFilename string

var parser = flags.NewParser(nil, flags.Default)

func init() {
	var err error
	home, err = homedir.Dir()
	util.CheckErr(err, "")

	todoDir = fmt.Sprintf("%s/.todo/", home)
	todoJsonFilename = fmt.Sprintf("%stodo", todoDir)
	archJsonFilename = fmt.Sprintf("%sarchive", todoDir)
	projJsonFilename = fmt.Sprintf("%sproject", todoDir)

	initTodoEnv()
	data.Todos, err = loadTodo(todoJsonFilename)
	util.CheckErr(err, "")
	data.ProjList, err = loadProject(projJsonFilename)
	util.CheckErr(err, "")
}

func notACommand(s string) bool {
	cmdList := []string{"add", "arch", "del", "do", "ls", "proj"}
	for _, v := range cmdList {
		if s == v {
			return false
		}
	}
	return true
}

func Run() {
	input := os.Args[1:]
	if opt := input[0]; notACommand(opt) {
		input = append(input, "")
		copy(input[1:], input[0:])
		input[0] = "add"
	}

	_, err := parser.ParseArgs(input)
	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}
