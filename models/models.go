// Package models is where all our models live
package models

type Registry struct {
	Todos Todos
}

func Register() *Registry {
	return &Registry{
		Todos: Todos{},
	}
}
