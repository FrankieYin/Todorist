package data

type ProjectList struct {
	Projects []*Project `json:"projects"`
}

func NewProjectList() *ProjectList {
	var projects = make([]*Project, 0)
	return &ProjectList{Projects: projects}
}
