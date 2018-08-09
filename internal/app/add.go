package app

import (
	"fmt"
	"os"
)

type AddCommand struct {
	Complete bool `short:"c" description:"Add and immediately complete a task"`
}

var add AddCommand

func init() {
	parser.AddCommand("add",
		"Add a todo",
		"The add command adds a todo to the todo-list. 'todo help add' to see more options.",
		&add)
}

func (cmd *AddCommand) Execute(args []string) error {
	if len(args) == 0 { // add cannot be called without an argument
		fmt.Println("No task specified, no task added.")
		fmt.Println("try 'todo help add' to see examples on how to add a task")
		os.Exit(0)
	}

	pTodoItem := parseTodo(args)
	todoList.AddTodo(pTodoItem)
	todoList.Save(todoJsonFilename)

	return nil
}