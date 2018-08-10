package data

type TodoItem struct {
	Task string `json:"task"`
	Due string `json:"due"`
	Project string `json:"project"`
	TimeCreated string `json:"time_created"`
	Done bool `json:"done"`
	Id int `json:"id"`	// does not change throughout the life time of the task
	ArchId int `json:"arch_id"`
	TimeArchived string `json:"time_archived"`
	Priority PriorityLevel `json:"priority"`
}
