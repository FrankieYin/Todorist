package app

import (
	"fmt"
	"os"
	"github.com/FrankieYin/todo/internal/data"
)

type DoCommand struct {
	All     bool `long:"all" description:"Complete all todos under the current focus."`
	Undo    bool `short:"u" long:"undo" description:"Un-complete todos specified by the ids."`
	Archive bool `short:"a" description:"Do and immediately archive a task."`
}

var do DoCommand

func init() {
	parser.AddCommand("do",
		"Complete todos specified by the ids.",
		"Complete todos specified by the ids. Use 'todo help do' to see more options",
		&do)
}

func (cmd *DoCommand) Execute(args []string) error {

	var ids []int
	currentFocus := data.ProjList.GetFocused()
	if do.All {
		for _, p := range currentFocus {
			ids = append(ids, p.Todos...)
		}
	} else {
		ids = parseId(args)
	}

	n := len(ids)
	if n == 0 {
		fmt.Println("No task Id specified under current focus, no task completed.")
		fmt.Println("try 'todo help do' to see examples on how to complete a task")
		os.Exit(0)
	}

	var err error
	undo := "Completed"
	if do.Undo {
		err = data.Todos.DoTodo(true, ids...)
		undo = "Un-completed"
	} else {
		err = data.Todos.DoTodo(false, ids...)
	}
	if err != nil {return err}

	if do.Archive {return arch.Execute(reverseId(ids...))}

	msg := "task"
	if n > 1 {msg = "tasks"}
	fmt.Printf("%s %d %s\n", undo, n, msg)

	return save()
}
