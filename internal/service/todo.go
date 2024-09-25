package service

import (
	"github.com/zuu-development/fullstack-examination-2024/internal/model"
	"github.com/zuu-development/fullstack-examination-2024/internal/repository"
)

type Todo interface {
	Create(task string) error
	Update(id int, task string, status model.TaskStatus) error
	Delete(id int) error
	Find(id int) (*model.Todo, error)
	FindAll() ([]*model.Todo, error)
}
type todo struct {
	todoRepository repository.Todo
}

func NewTodo(r repository.Todo) Todo {
	return &todo{r}
}
func (t *todo) Create(task string) error {
	todo := model.NewTodo(task)
	if err := t.todoRepository.Create(todo); err != nil {
		return err
	}
	return nil
}

func (t *todo) Update(id int, task string, status model.TaskStatus) error {
	todo := model.NewUpdateTodo(id, task, status)
	if err := t.todoRepository.Update(todo); err != nil {
		return err
	}
	return nil
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
