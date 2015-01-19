// Package models is where all our models live
package models

type registry struct {
	Todos Todos
}

func Register() *registry {
	return &registry{
		Todos: Todos{},
	}
}
