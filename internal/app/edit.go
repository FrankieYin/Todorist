package app

import (
	"github.com/FrankieYin/todo/internal/data"
	"github.com/FrankieYin/todo/internal/util"
		)

type EditCommand struct {
	Args struct{
		Id int
	} `positional-args:"yes" required:"yes"`
	Project bool `short:"p"`
	Plain bool `short:"T"` // suppress parsing, add the task as it is
	Due string `short:"d" long:"due" optional:"yes"`
}

var e EditCommand

func init()  {
	parser.AddCommand("e",
		"Edit information about a todo",
		"The edit command edits a todo. 'todo help e' to see more options.",
		&e)
}

/**
Usage:
e <id> <new_task>
e <id> -p <new_proj>
 */
func (cmd *EditCommand) Execute(args []string) error {
	id := e.Args.Id
	var err error
	pTodo, ok := data.Todos.Data[id]
	if !ok {
		return util.InvalidIdError{Id:id}
	}

	if e.Project {
		if err = pTodo.ChangeProject(args[0]); err != nil {return err}
	} else {
		if pTodo, err = parseTodo(args, e.Plain); err != nil {return err}
		pTodo.Id = id
		data.Todos.Data[id] = pTodo
	}

	return save()
}
