package v1

import (
	"github.com/gofiber/fiber/v2"
	todocontroller "github.com/heggies/todo-server/src/controller/v1/todo"
	"github.com/heggies/todo-server/src/usecase/v1/todo"
)

func StartHandler(c *fiber.App) (err error) {
	todoRepository, err := todo.NewRepository()
	if err != nil {
		return
	}

	todoService, err := todo.NewService(todoRepository)
	if err != nil {
		return
	}

	registerTodoController(c, todocontroller.NewController(todoService))

	return
}

func registerTodoController(c *fiber.App, controller *todocontroller.Controller) {
	c.Get("/todos", controller.Get)
	c.Post("/todos", controller.Create)
	c.Delete("/todos/:id", controller.Delete)
}
