package app

import (
	"strings"
)

func parse(input []string) *todoItem {
	return &todoItem{Task: strings.Join(input, " "), Done: false, Id: assignId()}
}

/**
 the principle is to assign the smallest available id the a newly created task
 */
func assignId() int {

	if len(todoList) == 0 {
		return 1
	}

	id := 1
	for _, pTodo := range todoList {
		if id == pTodo.Id {
			id++
		}
	}
	return id
}