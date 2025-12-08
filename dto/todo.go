package dto

type CreateTodoInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CategoryID  uint   `json:"category_id"`
}
