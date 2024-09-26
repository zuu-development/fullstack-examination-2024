package model

import "time"

// Todo is the model for the todo endpoint.
type Todo struct {
	ID        int `gorm:"primaryKey"`
	Task      string
	Status    TaskStatus
	CreatedAt time.Time `gorm:"<-:false"`
	UpdatedAt time.Time `gorm:"<-:false"`
}

// NewTodo returns a new instance of the todo model.
func NewTodo(task string) *Todo {
	return &Todo{
		Task:   task,
		Status: Created,
	}
}

// NewUpdateTodo returns a new instance of the todo model for updating.
func NewUpdateTodo(id int, task string, status TaskStatus) *Todo {
	return &Todo{
		ID:     id,
		Task:   task,
		Status: status,
	}
}

// TaskStatus is the status of the task.
type TaskStatus string

const (
	// Created is the status for a created task.
	Created = TaskStatus("created")
	// Processing is the status for a processing task.
	Processing = TaskStatus("processing")
	// Done is the status for a done task.
	Done = TaskStatus("done")
)

// TaskStatusMap is a map of task status.
var TaskStatusMap = map[TaskStatus]bool{
	Created:    true,
	Processing: true,
	Done:       true,
}
