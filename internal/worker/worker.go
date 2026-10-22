package worker

import (
	"fmt"
	"log"
	"time"
)

type Task struct {
	ID           string
	FileTask     string
	FilePriority bool
}

type TaskStatus struct {
	TaskID string
	Time   string
	Status bool
}

func (t *Task) worker(id string, taskQueue <-chan Task, done chan<- TaskStatus) {
	defer func() {
		done <- TaskStatus{
			TaskID: t.ID,
			Status: true,
			Time:   time.Now().Format("2006-01-02 15:04:05"),
		}
	}()

	for task := range taskQueue {
		if err := processTask(task); err != nil {
			// Handle error
			log.Fatal(err)
		}
	}

	fmt.Println(fmt.Sprintf("Worker %s finished processing task %s", id, t.ID))
}

func processTask(task Task) error {
	// Simulate task processing
	time.Sleep(2 * time.Second)
	fmt.Printf("Processing task %s with file %s and priority %v\n", task.ID, task.FileTask, task.FilePriority)
	return nil
}
