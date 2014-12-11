package models

// Todo is a todo item, it has a title and a completed status
type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"isCompleted"`
}

// Complete makes a Todo complete
func (todo *Todo) Complete() {
	todo.IsCompleted = true
}

// Uncomplete makes a Todo no longer complete
func (todo *Todo) Uncomplete() {
	todo.IsCompleted = false
}
