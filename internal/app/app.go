package app

import (
	"fmt"
	"os"
)

func HandleList(input []string) {
	fmt.Println("todo called with directory list")
}

func HandleAdd(input []string)  {
	if len(input) == 0 { // add cannot be called without an argument
		fmt.Println("No task specified, no task added.")
		fmt.Println("try 'todo help add' to see examples on how to add a task")
		os.Exit(0)
	}
}

func HandleDone(input []string)  {

}

func HandleProject(input []string)  {

}