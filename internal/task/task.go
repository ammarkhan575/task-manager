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

func NewTask(id int, title string, description string, priority Priority, dueDate time.Time, tags []string) *Task {
	return &Task{
		ID: id, 
		Title: title,
		Description: description,
		Priority: priority,
		Status: StatusTodo,
		CreatedAt: time.Now(),
		Tags: []string{},
	}
}