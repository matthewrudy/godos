// Package controllers is where our controllers live
//
// Each controller should have a handler
//
//   foosHandler(m martini.Router) {
//     m.Get("/", fooIndex)
//     m.Get("/:id", fooShow)
//     ...
//   }
//
package controllers

import (
	"github.com/go-martini/martini"
)

// Handler handles all routes for our app
func Handler(m martini.Router) {
	m.Group("/todos", todosHandler)
}
