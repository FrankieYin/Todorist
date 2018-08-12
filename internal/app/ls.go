package app

import (
	"fmt"
	"os"
	"github.com/FrankieYin/todo/internal/data"
)

type LsCommand struct {
	Verbose []bool `short:"v" long:"verbose" description:"Print the time at which a todo was added as well."`
	Project bool `short:"p" description:"List and group the todos into their projects. Todos with no project will be put into 'Inbox'."`
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

	if ls.Project {
		for _, p := range projList.Projects {
			listTodos(p)
		}
		return nil
	}

	listTodos(projList.Projects[0])
	return nil
}

func listTodos(p *data.Project) {
	fmt.Println(p.Name)
	for _, v := range p.Todos {
		pTodo, ok := todoList.Data[v]
		if ok {
			done := " "
			if pTodo.Done {done = "X"}
			fmt.Printf("%d\t[%s]\t%s\n", pTodo.Id, done, pTodo.Task)
		}
	}
	fmt.Println("")
}

