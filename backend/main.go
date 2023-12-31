package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	db := InitDatabase()

	app.Use(cors.New())

	// Setup routes
	app.Get("/todos", GetTodos(db))
	app.Post("/todos", CreateTodo(db))
	app.Put("/todos/:id", UpdateTodo(db))
	app.Delete("/todos/:id", DeleteTodo(db))

	app.Listen(":3000")
}
