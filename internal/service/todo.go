// Package service provides the business logic for the todo endpoint.
package service

import (
	"github.com/zuu-development/fullstack-examination-2024/internal/model"
	"github.com/zuu-development/fullstack-examination-2024/internal/repository"
)

// Todo is the service for the todo endpoint.
type Todo interface {
	Create(task string) (*model.Todo, error)
	Update(id int, task string, status model.Status) (*model.Todo, error)
	Delete(id int) error
	Find(id int) (*model.Todo, error)
	FindAll() ([]*model.Todo, error)
}

type todo struct {
	todoRepository repository.Todo
}

// NewTodo creates a new Todo service.
func NewTodo(r repository.Todo) Todo {
	return &todo{r}
}

func (t *todo) Create(task string) (*model.Todo, error) {
	todo := model.NewTodo(task)
	if err := t.todoRepository.Create(todo); err != nil {
		return nil, err
	}
	return todo, nil
}

func (t *todo) Update(id int, task string, status model.Status) (*model.Todo, error) {
	todo := model.NewUpdateTodo(id, task, status)
	// 現在の値を取得
	currentTodo, err := t.Find(id)
	if err != nil {
		return nil, err
	}
	// 空文字列の場合、現在の値を使用
	if todo.Task == "" {
		todo.Task = currentTodo.Task
	}
	if todo.Status == "" {
		todo.Status = currentTodo.Status
	}
	if err := t.todoRepository.Update(todo); err != nil {
		return nil, err
	}
	return todo, nil
}

func (t *todo) Delete(id int) error {
	if err := t.todoRepository.Delete(id); err != nil {
		return err
	}
	return nil
}

func (t *todo) Find(id int) (*model.Todo, error) {
	todo, err := t.todoRepository.Find(id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (t *todo) FindAll() ([]*model.Todo, error) {
	todo, err := t.todoRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return todo, nil
}
