package service

import "main/repository"

type TodoService struct {
	TodoRepo *repository.TodoRepository
}

func NewTodoService(todoRepo *repository.TodoRepository) *TodoService {
	return &TodoService{
		TodoRepo: todoRepo,
	}
}

func (s *TodoService) CreateItem() {}
