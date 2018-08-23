package data

import (
	"time"
	"github.com/FrankieYin/todo/internal/util"
)

type TodoItem struct {
	Task string `json:"task"`
	Due time.Time `json:"due"`
	Project string `json:"project"`
	TimeCreated string `json:"time_created"`
	Done bool `json:"done"`
	Id int `json:"id"`	// does not change throughout the life time of the task
	ArchId int `json:"arch_id"`
	TimeArchived string `json:"time_archived"`
	Priority PriorityLevel `json:"priority"` // by default is p4
}

func (t *TodoItem) IsOverDue() bool {
	return time.Now().After(t.Due)
}


func (t *TodoItem) ChangeProject(newProj string) error {
	if p := ProjList.GetProject(newProj); p != nil {
		p.AddTodo(t.Id)
		t.Project = newProj
	} else {
		return util.ProjectNotFound{Name:newProj}
	}
	return nil
}