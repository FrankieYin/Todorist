package app

import (
	"encoding/json"
	"os"

	"github.com/FrankieYin/Todorist/internal/util"
	"fmt"
)

type todoItem struct {
	Task string `json:"task"`
	Due string `json:"due"`
	Project string `json:"project"`
	TimeCreated string `json:"time_created"`
	Done bool `json:"done"`
	Id int `json:"id"`	// does not change throughout the life time of the task
}

func newTodoItem() {

}

func (item *todoItem) save()  {
	b, err := json.Marshal(item)
	util.CheckErr(err, "")

	f, err := os.OpenFile(jsonFilename, os.O_APPEND|os.O_WRONLY, 0644)
	util.CheckErr(err, "Error opening json file")

	defer f.Close()

	todo := fmt.Sprintf("%s\n", string(b))
	_, err = f.WriteString(todo)
	util.CheckErr(err, "Error writing json file")
}
