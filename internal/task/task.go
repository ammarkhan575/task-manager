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

// JSON struct tags control marshaling:
// `json:"id"`           → key name in JSON
// `json:"due,omitempty"` → omit field if zero value
// `json:"-"`            → always skip this field
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Priority    Priority  `json:"priority"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	DueDate     time.Time `json:"due_date,omitempty"`
	Tags        []string  `json:"tags"`
}

func NewTask(id int, title string, description string, priority Priority) *Task {
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

func (t *Task) AddTags(tags ...string) {
	t.Tags = append(t.Tags, tags...)
}
