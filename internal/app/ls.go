package app

import (
	"fmt"
	"os"
	"github.com/FrankieYin/todo/internal/data"
	"time"
)

type LsCommand struct {
	Verbose []bool `short:"v" long:"verbose" description:"Print the time at which a todo was added as well."`
	Project string `short:"p" long:"project" optional:"true" default:"true" description:"List and group the todos into their projects. Todos with no project will be put into 'Inbox'."`
	InboxOnly bool `short:"x"`
	All bool `long:"all" description:"list all todos grouped under projects"`
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
		if !ls.All {
			for _, p := range currentFocus {
				listProject(p, false)
			}
			return nil
		}
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
		pTodo := data.Todos.Data[id]
		done := " "
		projectName := ""
		if pTodo.Done {
			done = "X"
		}
		if showProject {
			if !(pTodo.Project == "") {
				projectName = fmt.Sprintf("%s: ", pTodo.Project)
			}
		}

		// show due date
		now := time.Now()
		today := time.Date(now.Year(), now.Month(), now.Day(), 0,0,0,0, now.Location())
		due := pTodo.Due
		dueDate := getDueString(today, due)
		fmt.Printf("%d\t[%s]   %-10s\t%s%s\n", pTodo.Id, done, dueDate, projectName, pTodo.Task)
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

func getDueString(today time.Time, due time.Time) string {
	if due.Equal(time.Time{}) {
		return ""
	}
	switch h := due.Sub(today).Hours(); h {
	case float64(-24), float64(0):
		return "Yesterday"
	case float64(24):
		return "Today"
	case float64(48):
		return "Tomorrow"
	}
	return due.Add(-time.Second).Format("Mon, Jan 2")
}

