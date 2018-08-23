package app

import (
	"fmt"
	"os"
	"github.com/FrankieYin/todo/internal/data"
	"github.com/FrankieYin/todo/internal/util"
	)

type AddCommand struct {
	Done bool `short:"c" description:"Add and immediately complete a task."`
	Archive bool `short:"a" description:"Add and immediately archive a task."`
	Priority int `short:"p" optional:"yes" default:"4"` // this priority flag if specified will override the priority parsed in task
	Plain bool `short:"T"` // suppress parsing, add the task as it is
	Due string `short:"d" long:"due" optional:"yes"`
}

var add AddCommand

func init() {
	parser.AddCommand("add",
		"Add a todo",
		"The add command adds a todo to the todo-list. 'todo help add' to see more options.",
		&add)
}

// add [-c | -a] <task>
// add [-p=<priority_level>] <task>
// <task> must be enclosed in double quotes ("")
func (cmd *AddCommand) Execute(args []string) error {
	if len(args) == 0 { // add cannot be called without an argument
		fmt.Println("No task specified, no task added.")
		fmt.Println("try 'todo help add' to see examples on how to add a task")
		os.Exit(0)
	}

	pTodoItem, err := parseTodo(args, add.Plain)
	if err != nil {return err}

	pTodoItem.Done = add.Done

	if level := getPriority(add.Priority); level == data.InvalidPriority {
		return util.InvalidPriorityLevel{Level:add.Priority}
	} else {
		pTodoItem.Priority = level
	}

	if add.Due != "" {
		if due, ok := getDue(add.Due); ok {
			pTodoItem.Due = due
		} else {
			msg := fmt.Sprintf("%s is not a valid date", add.Due)
			return util.InvalidArgument{Msg:msg}
		}
	}

	data.Todos.AddTodo(pTodoItem)
	if add.Archive { return arch.Execute(reverseId(pTodoItem.Id))}

	return save()
}

func getPriority(level int) data.PriorityLevel {
	switch level {
	case 4:
		return data.NotImportantNotUrgent
	case 3:
		return data.ImportantNotUrgent
	case 2:
		return data.NotImportantUrgent
	case 1:
		return data.ImportantUrgent
	default:
		return data.InvalidPriority
	}
}