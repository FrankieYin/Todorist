package app

import (
	"encoding/json"
	"io/ioutil"

	"github.com/FrankieYin/Todorist/internal/util"
)

type todoItem struct {
	Task        string `json:"task"`
	Due         string `json:"due"`
	Project     string `json:"project"`
	TimeCreated string `json:"time_created"`
}

func (item *todoItem) save()  {
	b, err := json.Marshal(item)
	util.CheckErr(err, "")

	err = ioutil.WriteFile(JsonFilename, b, 0644)
	util.CheckErr(err, "Error writing json file")
}
