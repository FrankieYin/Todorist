package data

type Project struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Todos []int `json:"todos"`
	Priority PriorityLevel `json:"priority"`
	TimeCreated string `json:"time_created"`
	TimeArchived string `json:"time_archived"`
	OnFocus bool `json:"on_focus"`
}

func (p *Project) AddTodo(ids ...int) {
	p.Todos = append(p.Todos, ids...)
}
