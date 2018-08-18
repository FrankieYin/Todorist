package data

type PriorityLevel int

const (
	InvalidPriority PriorityLevel = iota
	ImportantUrgent
	NotImportantUrgent
	ImportantNotUrgent
	NotImportantNotUrgent
)
