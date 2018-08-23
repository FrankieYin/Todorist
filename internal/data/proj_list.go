package data

import (
		"github.com/FrankieYin/todo/internal/util"
	)

type ProjectList struct {
	Projects []*Project `json:"projects"`
	Changed bool `json:"changed"`
}

func NewProjectList() *ProjectList {
	var projects = make([]*Project, 0)
	return &ProjectList{Projects: projects}
}

func (l *ProjectList) changed() {
	l.Changed = true
}

func (l *ProjectList) DeleteProject(name string) error {
	if i := l.indexOfProject(name); i != -1 {
		// delete all todos belonged to this project first
		Todos.DeleteTodo(l.Projects[i].Todos...)
		l.Projects = append(l.Projects[:i], l.Projects[i+1:]...)
		l.changed()
		return nil
	}
	return util.ProjectNotFound{Name:name}
}

func (l *ProjectList) RenameProject(oldName, newName string) error {
	if p := l.GetProject(oldName); p != nil {
		p.Name = newName
		l.changed()
		return nil
	}
	return util.ProjectNotFound{Name:oldName}
}

func (l *ProjectList) indexOfProject(name string) int {
	for i, p := range l.Projects {
		if p.Name == name {
			return i
		}
	}
	return -1
}

func (l *ProjectList) ContainsProject(name string) bool {
	return l.indexOfProject(name) != -1
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
	l.changed()
}

func (l *ProjectList) GetFocused() []*Project {
	currentFocus := make([]*Project, 0)
	for _, p := range l.Projects {
		if p.OnFocus {
			currentFocus = append(currentFocus, p)
		}
	}
	return currentFocus
}

func (l *ProjectList) ChangeFocus(names []string) error {
	for _, name := range names { // make sure all specified projects exist
		if !l.ContainsProject(name) {
			return util.ProjectNotFound{Name: name}
		}
	}

	for _, name := range names {
		l.GetProject(name).ChangeFocus()
	}
	l.changed()
	return nil
}

func (l *ProjectList) AddTodo(pTodo *TodoItem) error {
	if pTodo.Project == "" {return nil}
	l.GetProject(pTodo.Project).AddTodo(pTodo.Id)
	l.changed()
	return nil
}

func (l *ProjectList) DeleteTodo(pTodo *TodoItem) error {
	if pTodo.Project == "" {return nil}
	l.GetProject(pTodo.Project).DeleteTodo(pTodo.Id)
	l.changed()
	return nil
}

func (l *ProjectList) Save() error {
	if !l.Changed {return nil}
	return save(ProjList, projJsonFilename)
}
