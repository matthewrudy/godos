package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/matthewrudy/godos/controllers"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())

	// Handle all routes in controllers.Handler
	m.Group("", controllers.Handler)
	m.Run()
}
