package todo

import "github.com/gofiber/fiber/v2"

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (ctrl Controller) Get(c *fiber.Ctx) (err error) {
	c.JSON("Hello world!")

	return
}
