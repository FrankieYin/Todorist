package app

import (
	"fmt"
	"os"
	)

type AddCommand struct {
	Done bool `short:"c" description:"Add and immediately complete a task."`
	Archive bool `short:"a" description:"Add and immediately archive a task."`
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

	pTodoItem, err := parseTodo(args)
	if err != nil {return err}
	pTodoItem.Done = add.Done
	todoList.AddTodo(pTodoItem)
	if name := pTodoItem.Project; name != "" {
		projList.GetProject(name).AddTodo(pTodoItem.Id)
	}
	if add.Archive { return arch.Execute(reverseId(pTodoItem.Id))}

	if err = save(todoList, todoJsonFilename); err != nil {return err}
	return save(projList, projJsonFilename)
}