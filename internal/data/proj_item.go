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

func (p *Project) DeleteTodo(ids ...int) {
	for _, id := range ids {
		if i := p.IndexOfTodo(id); i != -1 {
			p.Todos = append(p.Todos[:i], p.Todos[i+1:]...)
		}
	}
}

func (p *Project) IndexOfTodo(id int) int {
	for i, t := range p.Todos {
		if t == id {
			return i
		}
	}
	return -1
}
