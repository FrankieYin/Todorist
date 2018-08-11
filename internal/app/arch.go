package app

import (
	"fmt"
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
	archList, err = loadTodo(archJsonFilename)
	if err != nil {return err}

	ids := parseId(args)
	archived, err := todoList.ArchTodo(len(archList.Data), ids...)
	if err != nil {return err}
	archList.Merge(archived)

	msg := "task"
	n := len(archived.Data)
	if n > 1 {msg = "tasks"}
	fmt.Printf("Archived %d %s\n", n, msg)

	if err = save(todoList, todoJsonFilename); err != nil { return err }
	if err = save(archList, archJsonFilename); err != nil { return err }

	return nil
}
