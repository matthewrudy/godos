package controllers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/matthewrudy/godos/models"
)

func todosHandler(m martini.Router) {
	m.Get("/", renderTodos(helloWorldTodo))
	m.Post("/complete", func() string {
		completeTodo(helloWorldTodo)
		return "OK"
	})
}

type todoList struct {
	Todos []*models.Todo `json:"todos"`
}

type renderer func(render.Render)

func renderTodos(todos ...*models.Todo) renderer {
	return func(r render.Render) {
		r.JSON(200, newTodoList(todos))
	}
}

func completeTodo(todo *models.Todo) *models.Todo {
	todo.Complete()
	return todo
}

func newTodoList(todos []*models.Todo) *todoList {
	return &todoList{
		Todos: todos,
	}
}

var helloWorldTodo = &models.Todo{
	Title:       "Hello world!",
	IsCompleted: false,
}
