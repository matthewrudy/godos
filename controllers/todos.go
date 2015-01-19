package controllers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/matthewrudy/godos/models"
)

func todosHandler(m martini.Router) {
	m.Get("", indexTodos)
	m.Get("/:id", showTodo)
	m.Post("/:id/complete", completeTodo)
}

func showTodo(r render.Render, p martini.Params) {
	todo := models.Register().Todos.FindByID(idParam(p))
	renderTodo(r, todo)
}

func completeTodo(r render.Render, p martini.Params) {
	todo := models.Register().Todos.FindByID(idParam(p))
	todo.Complete()
	renderTodo(r, todo)
}

func idParam(p martini.Params) string {
	return p["id"]
}

func indexTodos(r render.Render) {
	todos := models.Register().Todos.FindAll()
	renderTodos(r, todos)
}

func renderTodo(r render.Render, todo *models.Todo) {
	renderTodos(r, []*models.Todo{todo})
}

func renderTodos(r render.Render, todos []*models.Todo) {
	r.JSON(200, newTodoList(todos))
}

type todoList struct {
	Todos []*models.Todo `json:"todos"`
}

func newTodoList(todos []*models.Todo) *todoList {
	return &todoList{
		Todos: todos,
	}
}
