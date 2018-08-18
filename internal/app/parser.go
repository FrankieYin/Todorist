package app

import (
	"strings"
	"strconv"

	"github.com/FrankieYin/todo/internal/data"
	"github.com/FrankieYin/todo/internal/util"
	"time"
)

/**
 Usage for adding project
 research: finish building the PermissionFlowGraph class
 todolist: add support for parsing project names
 */
func parseTodo(input []string, suppressParsing bool) (*data.TodoItem, error) {
	project := ""
	if len(input) > 1 {return nil, util.TooManyArguments{Msg:"Too many arguments for adding a todo. " +
														"Did you enclose the todo in a double quote?"}}
	if suppressParsing {
		return &data.TodoItem{
			Task: input[0],
			Id: assignId(),
			Project:project,
			Priority:data.NotImportantNotUrgent}, nil
	}

	todo := strings.Split(input[0], " ")
	if arg := todo[0]; arg[len(arg)-1] == ':' { // a project is associated with the new task
		if project = arg[:len(arg)-1]; project == "" {
			return nil, util.InvalidArgument{Msg:"project name not specified."}
		}
		if !data.ProjList.ContainsProject(project) {
			return nil, util.ProjectNotFound{Name:project}
		}
		todo = todo[1:]
	}

	// in the future, can introduce nlp to parse dates
	// for now only support the following format:
	// due <date_identifier>
	parsedTodo := make([]string, len(todo))
	priority := data.NotImportantNotUrgent
	var dueDay time.Time
	for i, j := 0, 0; i < len(todo); i++ {
		s := todo[i]
		if strings.ToLower(s) == "due" {
			if due, ok := getDue(todo[i+1]); ok {
				dueDay = due
				i++
				continue
			}
		}
		switch s {
		case "p4":
			priority = data.NotImportantNotUrgent
		case "p3":
			priority = data.ImportantNotUrgent
		case "p2":
			priority = data.NotImportantUrgent
		case "p1":
			priority = data.ImportantUrgent
		default:
			parsedTodo[j] = s
			j++
		}
	}

	return &data.TodoItem{
		Task: strings.Join(parsedTodo, " "),
		Id:       assignId(),
		Project:  project,
		Priority: priority,
		Due:dueDay}, nil
}

func getDue(s string) (time.Time, bool) {
	now := time.Now()
	weekday := int(now.Weekday())
	var desiredDay int
	switch strings.ToLower(s) {
	case "tod", "today":
		desiredDay = weekday
	case "tom", "tomorrow":
		desiredDay = weekday + 1
	case "mon", "monday":
		desiredDay = 1
	case "tue", "tuesday":
		desiredDay = 2
	case "wed", "wednesday":
		desiredDay = 3
	case "thu", "thursday":
		desiredDay = 4
	case "fri", "friday":
		desiredDay = 5
	case "sat", "saturday":
		desiredDay = 6
	case "sun", "sunday":
		desiredDay = 0
	default:
		return time.Time{}, false
	}

	difference := desiredDay - weekday
	if difference < 0 {
		difference += 7
	}
	difference++

	return time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0,
		0,
		0,
		0,
		now.Location()).AddDate(0, 0, difference), true
}

func parseId(input []string) []int {
	var ids = make([]int , len(input))
	for i, idString := range input {
		id, err := strconv.Atoi(idString)
		util.CheckErr(err, "")
		ids[i] = id
	}
	return ids
}

func reverseId(ids ...int) []string {
	args := make([]string, len(ids))
	for i, id := range ids {
		args[i] = strconv.Itoa(id)
	}
	return args
}

/**
 the principle is to assign the smallest available id the a newly created task
 */
func assignId() int {

	if len(data.Todos.Data) == 0 {
		return 1
	}

	ids := make(map[int]bool)
	for k := range data.Todos.Data {
		ids[k] = true
	}

	var i int
	for i = range data.Todos.Order {
		_, ok := ids[i+1]
		if !ok {
			return i+1
		}
	}
	return i+2
}