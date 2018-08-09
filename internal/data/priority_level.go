package data

type PriorityLevel int

const (
	IMPORTANT_URGENT PriorityLevel = iota + 1
	NOT_IMPORTANT_URGENT
	IMPORTANT_NOT_URGENT
	NOT_IMPORTANT_NOT_URGENT
)
