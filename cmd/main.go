package main

import (
	"fmt"
	"time"

	"github.com/ghb-n/taskstream/internal/worker"
)

func main() {
	task := worker.Task{
		ID:           "1",
		FileTask:     "example.txt",
		FilePriority: true,
	}

	taskQueue := make(chan worker.Task, 1)
	taskTime := time.Now().Format("2006-01-02 15:04:05")
	done := make(chan worker.TaskStatus)
	taskStatus := true

	go worker.TaskStatus("1", taskTime, taskStatus)

	taskQueue <- task
	close(taskQueue)

	status := <-done
	fmt.Printf("Task %s completed with status: %s at %s\n", status.TaskID, status.Status, status.Time)

	fmt.Println("Всем привет!")
}
