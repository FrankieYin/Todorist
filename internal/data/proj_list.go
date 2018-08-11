package data

import (
	"fmt"

	"github.com/FrankieYin/todo/internal/util"
)

type ProjectList struct {
	Projects []*Project `json:"projects"`
}

func NewProjectList() *ProjectList {
	var projects = make([]*Project, 0)
	return &ProjectList{Projects: projects}
}

func (l *ProjectList) DeleteProject(name string) error {
	if i := l.IndexOfProject(name); i != -1 {
		l.Projects = append(l.Projects[:i], l.Projects[i+1:]...)
		return nil
	}
	msg := fmt.Sprintf("Project %s does not exist.", name)
	return util.ProjectNotFound{Msg:msg}
}

func (l *ProjectList) RenameProject(oldName, newName string) error {
	if p := l.GetProject(oldName); p != nil {
		p.Name = newName
		return nil
	}
	msg := fmt.Sprintf("Project %s does not exist.", oldName)
	return util.ProjectNotFound{Msg:msg}
}

func (l *ProjectList) IndexOfProject(name string) int {
	for i, p := range l.Projects {
		if p.Name == name {
			return i
		}
	}
	return -1
}
func (l *ProjectList) GetProject(name string) *Project {
	for _, p := range l.Projects {
		if p.Name == name {
			return p
		}
	}
	return nil
}
func (l *ProjectList) AddProject(project *Project) {
	l.Projects = append(l.Projects, project)
}
