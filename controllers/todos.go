package controllers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/matthewrudy/godos/models"
)

func todosHandler(m martini.Router) {
	m.Get("", indexTodos)
	m.Post("/complete", func() string {
		completeTodo(helloWorldTodo)
		return "OK"
	})
	m.Get("/:id", showTodo)
}

func showTodo(r render.Render, p martini.Params) {
	id := p["id"]
	renderTodos(r, randomTodo(id))
}

func indexTodos(r render.Render) {
	// 3 random todos
	renderTodos(r, helloWorldTodo, helloWorldTodo, helloWorldTodo)
}

type todoList struct {
	Todos []*models.Todo `json:"todos"`
}

type renderer func(render.Render)

func renderTodos(r render.Render, todos ...*models.Todo) {
	r.JSON(200, newTodoList(todos))
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

func randomTodo(id string) *models.Todo {
	return &models.Todo{
		Id:          id,
		Title:       "Todo task: " + id,
		IsCompleted: false,
	}
}

var helloWorldTodo = randomTodo("hello")
