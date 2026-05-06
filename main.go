package main

import (
	"fmt"
	"time"
	"github.com/ammarkhan575/task-manager/internal/task"
)

func main() {
	fmt.Println("Task Manager v0.1")
	t := task.NewTask(
		1,
		"Finish Go project",
		"Complete the task manager project in Go",
		task.PriorityHigh,
		time.Now().Add(48*time.Hour),
		[]string{"go", "project"},
	)
	fmt.Printf("Task created: %+v\n", t)
}
