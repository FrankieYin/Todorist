package app


import (
	"encoding/json"
	"os"
	"io/ioutil"

	"github.com/FrankieYin/Todorist/internal/util"
	"github.com/FrankieYin/Todorist/internal/data"
)

func save(list *data.TodoList, filename string) {
	// save todolist
	b, err := json.Marshal(list)
	util.CheckErr(err, "Unable to Marshal todolist")

	fTodo, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0644)
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
func loadTodo(filename string) *data.TodoList {
	b, err := ioutil.ReadFile(filename)
	util.CheckErr(err, "Error reading todo json file")

	var todos = new(data.TodoList)

	if len(b) == 0 { // empty json file
		return data.NewTodoList()
	}

	err = json.Unmarshal(b, todos)
	util.CheckErr(err, "Error Unmarshalling todo json file")

	return todos
}
