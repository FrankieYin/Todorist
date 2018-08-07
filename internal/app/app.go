package app

import (
	"fmt"
	"os"
	"encoding/json"
	"strconv"
	"io/ioutil"

	"github.com/mitchellh/go-homedir"
	"github.com/FrankieYin/Todorist/internal/util"
)

var home string
var todoJsonFilename string
var orderJsonFilename string
var todoDir string

var todoList map[int]*todoItem
var todoOrder []int

func Run(args []string) {
	initApp()
	execute(args)
}

func initApp() {
	var err error
	home, err = homedir.Dir()
	util.CheckErr(err, "")

	todoJsonFilename = fmt.Sprintf("%s/.todo/todo", home)
	orderJsonFilename = fmt.Sprintf("%s/.todo/order", home)
	todoDir = fmt.Sprintf("%s/.todo/", home)

	initTodoEnv()
	todoList = loadTodo()
	todoOrder = loadTodoOrder()
}

func execute(args []string) {
	command := args[0] // len(args) is guaranteed to be >= 1
	input := args[1:]
	// the above line will not give Out Of Bounds error because
	// we're slicing a slice and the bounds are 0 <= low <= high <= cap()

	switch command {
	case "list":
		handleList(input)
	case "add":
		handleAdd(input)
		save()
	case "done":
		handleDone(input)
		save()
	case "project":
		handleProject(input)
	}
}

func handleList(input []string) {

	if len(todoList) == 0 {
		fmt.Println("No task left undone!")
		fmt.Println("Use 'todo add' to add a new task.")
		os.Exit(0)
	}

	fmt.Println("All")
	for _, v := range todoOrder {
		pTodo, ok := todoList[v]
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

	pTodoItem := parse(input)
	todoList[pTodoItem.Id] = pTodoItem
	todoOrder = append(todoOrder, pTodoItem.Id)
}

func handleDone(input []string)  {
	n := len(input)
	if n == 0 {
		fmt.Println("No task Id specified, no task completed.")
		fmt.Println("try 'todo help done' to see examples on how to complete a task")
		os.Exit(0)
	}

	for _, idString := range input {
		id, err := strconv.Atoi(idString)
		util.CheckErr(err, "")
		pTodo, ok := todoList[id]
		if ok {
			pTodo.Done = true
		} else {
			fmt.Printf("todo done error: found no task with id %d\n", id)
			fmt.Println("Use 'todo list' first before completing a task")
			os.Exit(0)
		}
	}
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

	// save order list
	b, err = json.Marshal(todoOrder)
	util.CheckErr(err, "Unable to Marshal todo order list")

	fOrder, err := os.OpenFile(orderJsonFilename, os.O_WRONLY|os.O_TRUNC, 0644)
	util.CheckErr(err, "Error opening order json file")

	defer fOrder.Close()

	_, err = fOrder.Write(b)
	util.CheckErr(err, "Error writing order json file")
}

func initTodoEnv() {
	if _, err := os.Stat(todoDir); os.IsNotExist(err) {
		// create the directory
		err = os.Mkdir(todoDir, 0777)
		util.CheckErr(err, "Error creating directory /.todo")

		// initialise empty json files
		_, err = os.Create(todoJsonFilename)
		util.CheckErr(err, "failed to create json file")

		_, err = os.Create(orderJsonFilename)
		util.CheckErr(err, "failed to create json file")
	}
}

/**
 loads the json string into memory
 */
func loadTodo() map[int]*todoItem {
	b, err := ioutil.ReadFile(todoJsonFilename)
	util.CheckErr(err, "Error reading todo json file")

	var todoList = make(map[int]*todoItem)

	if len(b) == 0 { // empty json file
		return todoList
	}

	err = json.Unmarshal(b, &todoList)
	util.CheckErr(err, "Error Unmarshalling todo json file")

	return todoList
}

func loadTodoOrder() []int {
	b, err := ioutil.ReadFile(orderJsonFilename)
	util.CheckErr(err, "Error reading todo order json file")

	var todoOrder = make([]int, 0)

	if len(b) == 0 {
		return todoOrder
	}

	err = json.Unmarshal(b, &todoOrder)
	util.CheckErr(err, "Error Unmarshalling order json file")

	return todoOrder
}