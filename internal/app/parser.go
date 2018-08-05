package app

import "strings"

func parse(input []string) *todoItem {
	return &todoItem{Task:strings.Join(input, " ")}
}