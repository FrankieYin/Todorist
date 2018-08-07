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
		app.Run(args[1:])
	}
}

func printUsage() {
	fmt.Println("Thanks for using Todorist!")
}
