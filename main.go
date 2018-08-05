package main

import (
	"os"
	"fmt"

	"github.com/FrankieYin/Todorist/internal/app"
)

func main() {
	args := os.Args

	if len(args) <= 1 { // command invoked without a directory
		printUsage()
		os.Exit(1)
	} else {
		app.Init()
		execute(args[1:])
	}
}

func execute(args []string) {
	command := args[0] // len(args) is guaranteed to be >= 1
	input := args[1:]
	// the above line will not give Out Of Bounds error because
	// we're slicing a slice and the bounds are 0 <= low <= high <= cap()

	switch command {
	case "list":
		app.HandleList(input)
	case "add":
		app.HandleAdd(input)
	case "done":
		app.HandleDone(input)
	case "project":
		app.HandleProject(input)
	}
}

func printUsage() {
	fmt.Println("Thanks for using Todorist!")
}
