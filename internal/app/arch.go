package app

import (
	"fmt"
		"github.com/FrankieYin/todo/internal/util"
)

type ArchCommand struct {
}

var arch ArchCommand

func init() {
	parser.AddCommand("arch",
		"Archive todos specified by the ids.",
		"Archive all todos done by default.",
		&arch)
}

func (cmd *ArchCommand) Execute(args []string) error {
	archList = loadTodo(archJsonFilename)

	ids := parseId(args)
	archived, err := todoList.ArchTodo(len(archList.Data), ids...)
	util.CheckErr(err, "")
	archList.Merge(archived)

	msg := "task"
	n := len(archived.Data)
	if n > 1 {msg = "tasks"}
	fmt.Printf("Archived %d %s\n", n, msg)

	todoList.Save(todoJsonFilename)
	archList.Save(archJsonFilename)
	return nil
}
