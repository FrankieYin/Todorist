package app

import (
	"strings"
	"strconv"

	"github.com/FrankieYin/todo/internal/data"
	"github.com/FrankieYin/todo/internal/util"
)

func parseTodo(input []string) *data.TodoItem {
	return &data.TodoItem{Task: strings.Join(input, " "), Id: assignId()}
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