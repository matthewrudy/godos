package controllers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/matthewrudy/godos/models"
)

type Todos struct {
}

func (t *Todos) Handle(m martini.Router) {
	m.Get("", t.index)
	m.Get("/:id", t.show)
	m.Post("/:id/complete", t.complete)
}

func (t *Todos) show(r render.Render, p martini.Params) {
	todo := models.Register().Todos.FindByID(idParam(p))
	renderTodo(r, todo)
}

func (t *Todos) complete(r render.Render, p martini.Params) {
	todo := models.Register().Todos.FindByID(idParam(p))
	todo.Complete()
	renderTodo(r, todo)
}

func idParam(p martini.Params) string {
	return p["id"]
}

func (t *Todos) index(r render.Render) {
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
