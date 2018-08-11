package app

import (
		"fmt"
	"os"
)

type DelCommand struct {
}

var del DelCommand

func init() {
	parser.AddCommand("del",
		"Delete todos specified by the ids.",
		"Delete todos specified by the ids. If any of the id is invalid, no task will be deleted.",
		&del)
}

func (cmd *DelCommand) Execute(args []string) error {
	n := len(args)
	if n == 0 {
		fmt.Println("No task Id specified, no task deleted.")
		fmt.Println("try 'todo help del' to see examples on how to delete a task")
		os.Exit(0)
	}

	ids := parseId(args)

	if err := todoList.DeleteTodo(ids...); err != nil { return err}

	msg := "task"
	if n > 1 {msg = "tasks"}
	fmt.Printf("Deleted %d %s\n", n, msg)

	return todoList.Save(todoJsonFilename)
}
