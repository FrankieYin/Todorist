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
func loadTodo(filename string) (*data.TodoList, error){
	b, err := ioutil.ReadFile(filename)
	if err != nil {return nil, err}

	var todos = new(data.TodoList)

	if len(b) == 0 { // empty json file
		return data.NewTodoList(), nil
	}

	if err = json.Unmarshal(b, todos); err != nil {return nil, err}

	return todos, nil
}

func loadProject(filename string) (*data.ProjectList, error){
	b, err := ioutil.ReadFile(filename)
	if err != nil {return nil, err}

	var proj = new(data.ProjectList)

	if len(b) == 0 { // empty json file
		return data.NewProjectList(), nil
	}

	if err = json.Unmarshal(b, proj); err != nil {return nil, err}

	return proj, nil
}
