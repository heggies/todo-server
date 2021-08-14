package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heggies/todo-server/src/controller/v1/todo"
)

func StartHandler(c *fiber.App) {
	todoController(c, todo.NewController())
}

func todoController(c *fiber.App, controller *todo.Controller) {
	c.Get("/todo", controller.Get)
}
