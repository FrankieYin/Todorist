package app

import (
	"os"

	"github.com/jessevdk/go-flags"
		"github.com/FrankieYin/todo/internal/data"
)

var parser = flags.NewParser(nil, flags.Default)

func init() {
	data.InitTodoEnv()
}

func notACommand(s string) bool {
	cmdList := []string{"add", "arch", "del", "do", "e", "ls", "proj"}
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

func save() error {
	var err error
	if err = data.Todos.Save(); err != nil {return err}
	if err = data.ArchList.Save(); err != nil {return err}
	return data.ProjList.Save()
}
