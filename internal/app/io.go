package app


import (
	"encoding/json"
	"os"
	"io/ioutil"

	"github.com/FrankieYin/todo/internal/util"
	"github.com/FrankieYin/todo/internal/data"
)

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

func loadProject(filename string) *data.ProjectList {
	b, err := ioutil.ReadFile(filename)
	util.CheckErr(err, "Error reading todo json file")

	var proj = new(data.ProjectList)

	if len(b) == 0 { // empty json file
		return data.NewProjectList()
	}

	err = json.Unmarshal(b, proj)
	util.CheckErr(err, "Error Unmarshalling todo json file")

	return proj
}
