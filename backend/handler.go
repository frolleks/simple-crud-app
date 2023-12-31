package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetTodos(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var todos []Todo
		db.Find(&todos)
		return c.JSON(todos)
	}
}

func CreateTodo(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		todo := new(Todo)
		if err := c.BodyParser(todo); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		db.Create(&todo)
		return c.JSON(todo)
	}
}

type UpdateTodoInput struct {
	Title     *string `json:"title"`
	Completed *bool   `json:"completed"`
}

func UpdateTodo(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var todo Todo

		// Find the todo item by ID
		if err := db.First(&todo, id).Error; err != nil {
			return c.Status(404).SendString("Todo not found")
		}

		var input UpdateTodoInput
		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		// Update fields in todo if they are present in input
		if input.Title != nil {
			todo.Title = *input.Title
		}
		if input.Completed != nil {
			todo.Completed = *input.Completed
		}

		db.Save(&todo)

		return c.JSON(todo)
	}
}

func DeleteTodo(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var todo Todo

		// Find the todo by ID and delete it
		if err := db.Where("id = ?", id).Delete(&todo).Error; err != nil {
			return c.Status(500).SendString("Error deleting todo")
		}

		return c.SendStatus(204) // No Content
	}
}
