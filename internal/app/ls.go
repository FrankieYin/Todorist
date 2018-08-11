package app

import (
	"fmt"
	"os"
)

type LsCommand struct {
	Verbose []bool `short:"v" long:"verbose" description:"Print the time at which a todo was added as well."`
}

var ls LsCommand

func init() {
	parser.AddCommand("ls",
		"List todos.",
		"List all todos by default. Use 'todo help ls' to see how to filter and group todos.",
		&ls)
}

func (cmd *LsCommand) Execute(args []string) error {
	if len(todoList.Data) == 0 {
		fmt.Println("No todo left undone!")
		fmt.Println("Use 'todo add' to add a new task.")
		os.Exit(0)
	}

	fmt.Println("All")
	for _, v := range todoList.Order {
		pTodo, ok := todoList.Data[v]
		if ok {
			done := " "
			if pTodo.Done {done = "X"}
			fmt.Printf("%d\t[%s]\t%s\n", pTodo.Id, done, pTodo.Task)
		}
	}
	return nil
}

