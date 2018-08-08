package app

import (
	"fmt"
	"os"
			"github.com/mitchellh/go-homedir"
	"github.com/FrankieYin/Todorist/internal/util"
	"github.com/FrankieYin/Todorist/internal/data"
)

var home string
var todoJsonFilename string
var archJsonFilename string
var todoDir string

var todoList *data.TodoList
var archList *data.TodoList

func Run(args []string) {
	initApp()
	execute(args)
}

func initApp() {
	var err error
	home, err = homedir.Dir()
	util.CheckErr(err, "")

	todoDir = fmt.Sprintf("%s/.todo/", home)
	todoJsonFilename = fmt.Sprintf("%stodo", todoDir)
	archJsonFilename = fmt.Sprintf("%sarchive", todoDir)

	initTodoEnv()
	todoList = loadTodo(todoJsonFilename)
}

func execute(args []string) {
	command := args[0] // len(args) is guaranteed to be >= 1
	input := args[1:]
	// the above line will not give Out Of Bounds error because
	// we're slicing a slice and the bounds are 0 <= low <= high <= cap()

	switch command {
	case "ls":
		handleList(input)
	case "add":
		handleAdd(input)
		save(todoList, todoJsonFilename)
	case "done":
		handleDone(input)
		save(todoList, todoJsonFilename)
	case "proj":
		handleProject(input)
	case "del":
		handleDel(input)
		save(todoList, todoJsonFilename)
	case "arch":
		handleArch(input)
		save(todoList, todoJsonFilename)
		save(archList, archJsonFilename)
	default:
		fmt.Printf("todo has no command named '%s'\n", command) // todo implement command fuzzing
	}
}

func handleArch(input []string) {

	archList = loadTodo(archJsonFilename)

	ids := parseId(input)
	archived, err := todoList.ArchTodo(len(archList.Data), ids...)

	util.CheckErr(err, "")
	archList.Merge(archived)

	msg := "task"
	n := len(archived.Data)
	if n > 1 {msg = "tasks"}
	fmt.Printf("Archived %d %s\n", n, msg)
}

/**
 delete the task(s) specified by the ids. If any of the id is invalid, no task will be deleted.
 */
func handleDel(input []string) {
	n := len(input)
	if n == 0 {
		fmt.Println("No task Id specified, no task deleted.")
		fmt.Println("try 'todo help del' to see examples on how to delete a task")
		os.Exit(0)
	}

	ids := parseId(input)

	err := todoList.DeleteTodo(ids...)
	util.CheckErr(err, "")

	msg := "task"
	if n > 1 {msg = "tasks"}
	fmt.Printf("Deleted %d %s\n", n, msg)
}

func handleList(input []string) {

	if len(todoList.Data) == 0 {
		fmt.Println("No task left undone!")
		fmt.Println("Use 'todo add' to add a new task.")
		os.Exit(0)
	}

	fmt.Println("All")
	for _, v := range todoList.Order {
		pTodo, ok := todoList.Data[v]
		if ok {
			done := " "
			if pTodo.Done {done = "X"}
			fmt.Printf("%d\t[%s]\t%s\n", pTodo.Id, done, pTodo.Task)
		}
	}
}

/**
 usage:
 add finish todorist add functionality due today
 */
func handleAdd(input []string)  {
	if len(input) == 0 { // add cannot be called without an argument
		fmt.Println("No task specified, no task added.")
		fmt.Println("try 'todo help add' to see examples on how to add a task")
		os.Exit(0)
	}

	pTodoItem := parseTodo(input)
	todoList.AddTodo(pTodoItem)
}

func handleDone(input []string)  {
	n := len(input)
	if n == 0 {
		fmt.Println("No task Id specified, no task completed.")
		fmt.Println("try 'todo help done' to see examples on how to complete a task")
		os.Exit(0)
	}

	ids := parseId(input)

	err := todoList.DoneTodo(ids...)
	util.CheckErr(err, "")

	msg := "task"
	if n > 1 {msg = "tasks"}
	fmt.Printf("Completed %d %s\n", n, msg)
}

func handleProject(input []string)  {
	fmt.Println("todo called with directory project")
}