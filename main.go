package main

import (
	"fmt"
	"time"

	"github.com/ammarkhan575/task-manager/internal/task"
)

func addSampleTask(store task.StoreInterface) {
	t := task.NewTask(
		store.NextID(),
		"Finish Go project",
		"Complete the task manager project in Go",
		task.PriorityHigh,
		time.Now().Add(48*time.Hour),
		[]string{"go", "project"},
	)

	store.AddTask(t)
	fmt.Printf("Task created: %+v\n", t)
}

func main() {
	fmt.Println("Task Manager v0.1")
	var taskStore task.StoreInterface = task.NewStore()
	addSampleTask(taskStore)
	tasks := taskStore.GetAll()
	fmt.Printf("All tasks: %+v\n", tasks)
	// for(_, task := range tasks) {
	// 	fmt.Printf("Task: %+v\n", *task)
	// }
}
