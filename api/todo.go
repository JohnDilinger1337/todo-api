package api

import "main/controller"

type TodoAPI struct {
	TodoCtrl controller.TodoController
}

func NewTodoAPI(todoCtrl *controller.TodoController) *TodoAPI {
	return &TodoAPI{
		TodoCtrl: *todoCtrl,
	}
}
