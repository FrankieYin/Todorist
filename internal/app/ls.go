package app

import (
	"fmt"
	"os"
	"github.com/FrankieYin/todo/internal/data"
)

type LsCommand struct {
	Verbose []bool `short:"v" long:"verbose" description:"Print the time at which a todo was added as well."`
	Project string `short:"p" long:"project" optional:"true" default:"true" description:"List and group the todos into their projects. Todos with no project will be put into 'Inbox'."`
	InboxOnly bool `short:"x"`
}

var ls LsCommand

func init() {
	parser.AddCommand("ls",
		"List todos.",
		"List all todos by default. Use 'todo help ls' to see how to filter and group todos.",
		&ls)
}

func (cmd *LsCommand) Execute(args []string) error {
	if len(data.Todos.Data) == 0 {
		fmt.Println("No todo left undone!")
		fmt.Println("Use 'todo add' to add a new task.")
		os.Exit(0)
	}

	currentFocus := data.ProjList.GetFocused()

	if len(currentFocus) != 0 { // list only the todos in current focus
		for _, p := range currentFocus {
			listProject(p, false)
		}
		return nil
	}

	listAll(ls.InboxOnly)

	if ls.Project == "true" {
		for _, p := range data.ProjList.Projects {
			listProject(p, false)
		}
		return nil
	}

	return nil
}

func listProject(p *data.Project, showProject bool) {
	fmt.Println(p.Name)
	listTodos(p.Todos, false)
}

func listTodos(ids []int, showProject bool) {
	for _, id := range ids {
		pTodo, ok := data.Todos.Data[id]
		if ok {
			done := " "
			projectName := ""
			if pTodo.Done {done = "X"}
			if showProject {projectName = fmt.Sprintf("%s: ", pTodo.Project)}
			fmt.Printf("%d\t[%s]\t%s%s\n", pTodo.Id, done, projectName, pTodo.Task)
		}
	}
	fmt.Println("")
}

func listAll(inboxOnly bool) {
	fmt.Println("Inbox")
	var inboxList []int
	if inboxOnly {
		inboxList = make([]int, 0)
		for _, id := range data.Todos.Order {
			if pTodo := data.Todos.Data[id]; pTodo.Project == "" {
				inboxList = append(inboxList, id)
			}
		}
	} else {
		inboxList = data.Todos.Order
	}

	listTodos(inboxList, true)
}

