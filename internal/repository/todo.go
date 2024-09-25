package repository

import (
	"github.com/zuu-development/fullstack-examination-2024/internal/model"
	"gorm.io/gorm"
)

type Todo interface {
	Create(t *model.Todo) error
	Delete(id int) error
	Update(t *model.Todo) error
	Find(id int) (*model.Todo, error)
	FindAll() ([]*model.Todo, error)
}

type todo struct {
	db *gorm.DB
}

func NewTodo(db *gorm.DB) Todo {
	return &todo{
		db: db,
	}
}

func (td *todo) Create(t *model.Todo) error {
	if err := td.db.Create(t).Error; err != nil {
		return err
	}
	return nil
}

func (td *todo) Update(t *model.Todo) error {
	if err := td.db.Save(t).Error; err != nil {
		return err
	}
	return nil
}

func (td *todo) Delete(id int) error {
	if err := td.db.Where("id = ?", id).Delete(&model.Todo{}).Error; err != nil {
		return err
	}
	return nil
}

func (td *todo) Find(id int) (*model.Todo, error) {
	var todo *model.Todo
	err := td.db.Where("id = ?", id).Take(&todo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return todo, nil
}
func (td *todo) FindAll() ([]*model.Todo, error) {
	var todos []*model.Todo
	err := td.db.Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}
