package app

import (
	"fmt"
	"os"
	"encoding/json"
		"io/ioutil"

	"github.com/mitchellh/go-homedir"
	"github.com/FrankieYin/Todorist/internal/util"
	"github.com/FrankieYin/Todorist/internal/data"
)

var home string
var todoJsonFilename string
var archJsonFilename string
var todoDir string

var todoList *data.TodoList

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
	todoList = loadTodo()
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
		save()
	case "done":
		handleDone(input)
		save()
	case "proj":
		handleProject(input)
	case "del":
		handleDel(input)
		save()
	case "arch":
		handleArch(input)
		save()
	default:
		fmt.Printf("todo has no command named '%s'\n", command) // todo implement command fuzzing
	}
}

func handleArch(input []string) {
	n := len(input)
	if n == 0 { // when no id specified, archive all tasks done
		for k, pTodo := range todoList.Data {
			if pTodo.Done {
				delete(todoList.Data, k)
				// todo
			}
		}
	}
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

func save() {
	// save todolist
	b, err := json.Marshal(todoList)
	util.CheckErr(err, "Unable to Marshal todolist")

	fTodo, err := os.OpenFile(todoJsonFilename, os.O_WRONLY|os.O_TRUNC, 0644)
	util.CheckErr(err, "Error opening todo json file")

	defer fTodo.Close()

	_, err = fTodo.Write(b)
	util.CheckErr(err, "Error writing todo json file")
}

func initTodoEnv() {
	if _, err := os.Stat(todoDir); os.IsNotExist(err) {
		// create the directory
		err = os.Mkdir(todoDir, 0777)
		util.CheckErr(err, "Error creating directory /.todo")

		// initialise empty json files
		_, err = os.Create(todoJsonFilename)
		util.CheckErr(err, "failed to create json file")

		_, err = os.Create(archJsonFilename)
		util.CheckErr(err, "failed to create json file")
	}
}

/**
 loads the json string into memory
 */
func loadTodo() *data.TodoList {
	b, err := ioutil.ReadFile(todoJsonFilename)
	util.CheckErr(err, "Error reading todo json file")

	var todos = new(data.TodoList)

	if len(b) == 0 { // empty json file
		return data.NewTodoList()
	}

	err = json.Unmarshal(b, todos)
	util.CheckErr(err, "Error Unmarshalling todo json file")

	return todos
}