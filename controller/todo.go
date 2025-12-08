package controller

import "main/service"

type TodoController struct {
	TodoSvc *service.TodoService
}

func NewTodoController(todoSvc *service.TodoService) *TodoController {
	return &TodoController{TodoSvc: todoSvc}
}
