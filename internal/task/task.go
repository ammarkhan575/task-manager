package task

import "time"

type Priority int

const (
	PriorityLow Priority = iota
	PriorityMedium
	PriorityHigh
)

type Status string

const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in_progress"
	StatusDone       Status = "done"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Priority    Priority
	Status      Status
	CreatedAt   time.Time
	DueDate     time.Time
	Tags        []string
}
