package app

import (
	"fmt"
	"os"
	"log"
		"bufio"

	"github.com/mitchellh/go-homedir"
	"github.com/FrankieYin/Todorist/internal/util"
	"encoding/json"
)

var home string
var jsonFilename string
var todoList []*todoItem

func Init() {
	var err error
	home, err = homedir.Dir()
	util.CheckErr(err, "")

	jsonFilename = fmt.Sprintf("%s/.todo/todo", home)
	todoList = loadTodo()
}

func HandleList(input []string) {
	fmt.Println("All")
	for _, pTodo := range todoList {
		fmt.Printf("[ ]\t%s\n", pTodo.Task)
	}
}

/**
 usage:
 add finish todorist add functionality due today
 */
func HandleAdd(input []string)  {
	if len(input) == 0 { // add cannot be called without an argument
		fmt.Println("No task specified, no task added.")
		fmt.Println("try 'todo help add' to see examples on how to add a task")
		os.Exit(0)
	}

	pTodoItem := parse(input)
	pTodoItem.save()
}

func HandleDone(input []string)  {
	fmt.Println("todo called with directory done")
}

func HandleProject(input []string)  {
	fmt.Println("todo called with directory project")
}

/**
 loads the json string into memory
 */
func loadTodo() []*todoItem {
	f, err := os.OpenFile(jsonFilename, os.O_RDONLY, 0444)
	if err != nil {
		switch {
		case os.IsNotExist(err):
			// create the directory
			dir := fmt.Sprintf("%s/.todo/", home)
			err = os.Mkdir(dir, 0777)
			util.CheckErr(err, "Error creating directory /.todo")

			// initialise an empty json file
			f, err := os.Create(jsonFilename)
			util.CheckErr(err, "failed to create json file")

			defer f.Close()

			return loadTodo()
		case os.IsPermission(err):
			log.Fatal("file read permission denied.")
		}
	}

	defer f.Close()

	var pTodo *todoItem
	var todoList = make([]*todoItem, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		pTodo = new(todoItem)
		jsonString := scanner.Text()
		err = json.Unmarshal([]byte(jsonString), pTodo)

		util.CheckErr(err, "")

		todoList = append(todoList, pTodo)
	}
	util.CheckErr(scanner.Err(), "An error occurred during scanning json file")

	return todoList
}