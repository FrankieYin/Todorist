package data

type Project struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Priority PriorityLevel `json:"priority"`
	TimeCreated string `json:"time_created"`
	TimeArchived string `json:"time_archived"`
	OnFocus bool `json:"on_focus"`
}
