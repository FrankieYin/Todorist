package data


import (
	"encoding/json"
	"os"
	"io/ioutil"

	"github.com/FrankieYin/todo/internal/util"
	"github.com/mitchellh/go-homedir"
	"fmt"
)

var home string
var todoDir string
var todoJsonFilename string
var archJsonFilename string
var projJsonFilename string

func InitTodoEnv() {
	var err error
	home, err = homedir.Dir()
	util.CheckErr(err, "")

	todoDir = fmt.Sprintf("%s/.todo/", home)
	todoJsonFilename = fmt.Sprintf("%stodo", todoDir)
	archJsonFilename = fmt.Sprintf("%sarchive", todoDir)
	projJsonFilename = fmt.Sprintf("%sproject", todoDir)

	if _, err := os.Stat(todoDir); os.IsNotExist(err) {
		// create the directory
		err = os.Mkdir(todoDir, 0777)
		util.CheckErr(err, "Error creating directory /.todo")

		// initialise empty json files
		_, err = os.Create(todoJsonFilename)
		util.CheckErr(err, "failed to create json file")

		_, err = os.Create(archJsonFilename)
		util.CheckErr(err, "failed to create json file")

		_, err = os.Create(projJsonFilename)
		util.CheckErr(err, "failed to create json file")
	}

	Todos, err = loadTodo(todoJsonFilename)
	util.CheckErr(err, "")
	ArchList, err = loadTodo(archJsonFilename)
	util.CheckErr(err, "")
	ProjList, err = loadProject(projJsonFilename)
	util.CheckErr(err, "")
}

/**
 loads the json string into memory
 */
func loadTodo(filename string) (*TodoList, error){
	b, err := ioutil.ReadFile(filename)
	if err != nil {return nil, err}

	var todos = new(TodoList)

	if len(b) == 0 { // empty json file
		return NewTodoList(), nil
	}

	if err = json.Unmarshal(b, todos); err != nil {return nil, err}

	return todos, nil
}

func loadProject(filename string) (*ProjectList, error){
	b, err := ioutil.ReadFile(filename)
	if err != nil {return nil, err}

	var proj = new(ProjectList)

	if len(b) == 0 { // empty json file
		return NewProjectList(), nil
	}

	if err = json.Unmarshal(b, proj); err != nil {return nil, err}

	return proj, nil
}

func save(v interface{}, filename string) error {
	var b []byte
	var err error
	var f *os.File
	if b, err = json.Marshal(v); err != nil { return err}

	if f, err = os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644); err != nil { return err}

	defer f.Close()

	if _, err = f.Write(b); err != nil { return err}

	return nil
}
