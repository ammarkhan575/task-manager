package task

import (
	"fmt"
	"time"
)

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
		ID:          id,
		Title:       title,
		Description: description,
		Priority:    priority,
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		Tags:        []string{},
	}
}

func (t *Task) String() string {
	due := "no due date"
	if !t.DueDate.IsZero() {
		due = t.DueDate.Format("2006-01-02 15:04")
	}
	return fmt.Sprintf("[%d] | %s | %s  | %s | due: %s", t.ID, t.Title, t.Priority, t.Status, due)
}

func (p Priority) String() string {
	// switch p {
	// case PriorityLow:
	// 	return "Low"
	// case PriorityMedium:
	// 	return "Medium"
	// case PriorityHigh:
	// 	return "High"
	// default:
	// 	return "Unknown"
	// }
	return [...]string{"Low", "Medium", "High"}[p]
}

func (s Status) String() string { return string(s) }

func (t *Task) Start() { t.Status = StatusInProgress }

func (t *Task) Complete() { t.Status = StatusDone }
