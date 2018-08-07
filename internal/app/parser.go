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

	ids := make(map[int]bool)
	for k := range todoList {
		ids[k] = true
	}

	var i int
	for i = range todoOrder {
		_, ok := ids[i+1]
		if !ok {
			return i+1
		}
	}
	return i+2
}