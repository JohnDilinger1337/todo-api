package repository

import (
	"errors"
	"main/database/model"

	"gorm.io/gorm"
)

type TodoRepository struct {
	DB *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{DB: db}
}

func (r *TodoRepository) CreateCategory(category *model.TodoCategory) error {
	var exist model.TodoCategory
	if err := r.DB.Where("name = ?", category.Name).First(&exist); err != nil {
		return errors.New("something went wrong while trying to create category")
	}

	if err := r.DB.Create(category); err != nil {
		return errors.New("error while trying to create category")
	}

	return nil
}

func (r *TodoRepository) CreateTodoItem(item *model.TodoItem) error {
	var exist model.TodoItem
	if err := r.DB.Where("name = ? AND category_id = ?", item.Name, item.CategoryID).First(&exist); err != nil {
		return errors.New("something went wrong while trying to create todo item")
	}

	if err := r.DB.Create(item); err != nil {
		return errors.New("error while trying to create todo item")
	}

	return nil
}
