package app

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"

	"github.com/mitchellh/go-homedir"
	"github.com/FrankieYin/Todorist/internal/util"
)

var HOME string
var JsonFilename string

func Init() {
	var err error
	HOME, err = homedir.Dir()
	util.CheckErr(err, "")

	JsonFilename = fmt.Sprintf("%s/.todo/todo.json", HOME)
	loadTodo(loadFile(JsonFilename))
}

func HandleList(input []string) {
	fmt.Println("todo called with directory list")
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
 loads the json file
 */
func loadFile(filename string) []byte {
	b, err := ioutil.ReadFile(filename)
	if err != nil { // this will happen only when the command is called for the first time
		switch {
		case os.IsNotExist(err):
			// create the directory
			dir := fmt.Sprintf("%s/.todo/", HOME)
			err = os.Mkdir(dir, 0777)
			util.CheckErr(err, "")

			// initialise an empty todo.json file
			f, err := os.Create(filename)
			util.CheckErr(err, "failed to create json file")

			defer f.Close()

			return loadFile(filename)
		case os.IsPermission(err):
			log.Fatal("file read permission denied.")
		}
	}
	return b
}

/**
 loads the json string into memory
 */
func loadTodo(b []byte) {
}