package app

import (
	"fmt"
	"github.com/FrankieYin/todo/internal/data"
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
	var err error

	ids := parseId(args)
	if err = data.Todos.ArchTodo(ids...); err != nil {return err}

	msg := "task"
	n := len(ids)
	if n > 1 {msg = "tasks"}
	fmt.Printf("Archived %d %s\n", n, msg)

	return save()
}
