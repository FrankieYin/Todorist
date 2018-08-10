package app

import (
	"fmt"
	"os"
	)

type DoneCommand struct {
	All bool `long:"all" description:"Complete all todos under the current focus."`
	Undone bool `short:"u" long:"undo" description:"Uncomplete todos specified by the ids"`
}

var done DoneCommand

func init() {
	parser.AddCommand("done",
		"Complete todos specified by the ids.",
		"Complete todos specified by the ids. Use 'todo help done' to see more options",
		&done)
}

func (cmd *DoneCommand) Execute(args []string) error {
	n := len(args)
	if n == 0 {
		fmt.Println("No task Id specified, no task completed.")
		fmt.Println("try 'todo help done' to see examples on how to complete a task")
		os.Exit(0)
	}

	ids := parseId(args)

	if err := todoList.DoneTodo(ids...); err != nil {return err}

	msg := "task"
	if n > 1 {msg = "tasks"}
	fmt.Printf("Completed %d %s\n", n, msg)

	return todoList.Save(todoJsonFilename)
}
