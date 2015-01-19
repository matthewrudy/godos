package models

import (
	"fmt"
)

// Todo is a todo item, it has a title and a completed status
type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"isCompleted"`
}

// String returns a human readable string
func (todo *Todo) String() string {
	return fmt.Sprintf("Todo#%s", todo.ID)
}

// FindByID finds a single Todo by its ID
func FindByID(id string) *Todo {
	return makeTodo(id)
}

// FindAll finds all Todos
func FindAll() []*Todo {
	return []*Todo{
		makeTodo("alpha"),
		makeTodo("beta"),
		makeTodo("gamma"),
	}
}

// Complete makes a Todo complete
func (todo *Todo) Complete() {
	todo.IsCompleted = true
}

// Uncomplete makes a Todo no longer complete
func (todo *Todo) Uncomplete() {
	todo.IsCompleted = false
}

func makeTodo(id string) *Todo {
	return &Todo{
		ID:          id,
		Title:       "Todo task: " + id,
		IsCompleted: false,
	}
}
