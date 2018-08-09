package app

import (
	"fmt"
	"os"

	"github.com/FrankieYin/todo/internal/data"
	"github.com/jessevdk/go-flags"
	"github.com/mitchellh/go-homedir"
	"github.com/FrankieYin/todo/internal/util"
)

var home string
var todoDir string
var todoJsonFilename string
var archJsonFilename string
var projJsonFilename string

var todoList *data.TodoList
var archList *data.TodoList

var projList *data.ProjectList

var parser = flags.NewParser(nil, flags.Default)

func init() {
	var err error
	home, err = homedir.Dir()
	util.CheckErr(err, "")

	todoDir = fmt.Sprintf("%s/.todo/", home)
	todoJsonFilename = fmt.Sprintf("%stodo", todoDir)
	archJsonFilename = fmt.Sprintf("%sarchive", todoDir)

	initTodoEnv()
	todoList = loadTodo(todoJsonFilename)
}

func Run() {
	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}