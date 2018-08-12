package app

import (
	"strings"
	"strconv"

	"github.com/FrankieYin/todo/internal/data"
	"github.com/FrankieYin/todo/internal/util"
)

/**
 Usage for adding project
 research: finish building the PermissionFlowGraph class
 todolist: add support for parsing project names
 */
func parseTodo(input []string) (*data.TodoItem, error) {
	project := ""
	if arg := input[0]; arg[len(arg)-1] == ':' { // a project is associated with the new task
		if project = arg[:len(arg)-1]; project == "" {
			return nil, util.InvalidArgument{Msg:"project name not specified."}
		}
		if !projList.ContainsProject(project) {
			return nil, util.ProjectNotFound{Name:project}
		}
	}

	return &data.TodoItem{Task: strings.Join(input, " "),
						  Id: assignId(),
	 					  Project:project}, nil
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

	if len(todoList.Data) == 0 {
		return 1
	}

	ids := make(map[int]bool)
	for k := range todoList.Data {
		ids[k] = true
	}

	var i int
	for i = range todoList.Order {
		_, ok := ids[i+1]
		if !ok {
			return i+1
		}
	}
	return i+2
}