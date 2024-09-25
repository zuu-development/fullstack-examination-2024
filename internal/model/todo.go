package model

import "time"

type Todo struct {
	ID        int `gorm:"primaryKey"`
	Task      string
	Status    TaskStatus
	CreatedAt time.Time `gorm:"<-:false"`
	UpdatedAt time.Time `gorm:"<-:false"`
}

func NewTodo(task string) *Todo {
	return &Todo{
		Task:   task,
		Status: Created,
	}
}
func NewUpdateTodo(id int, task string, status TaskStatus) *Todo {
	return &Todo{
		ID:     id,
		Task:   task,
		Status: status,
	}
}

type TaskStatus string

const (
	Created    = TaskStatus("created")
	Processing = TaskStatus("processing")
	Done       = TaskStatus("done")
)

var TaskStatusMap = map[TaskStatus]bool{
	Created:    true,
	Processing: true,
	Done:       true,
}
