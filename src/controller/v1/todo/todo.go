package todo

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/heggies/todo-server/src/controller/v1/todo/presenter"
	"github.com/heggies/todo-server/src/entity/v1/todo"
	usecase "github.com/heggies/todo-server/src/usecase/v1/todo"
	"github.com/jinzhu/copier"
)

type Controller struct {
	s usecase.Service
}

func NewController(s *usecase.Service) *Controller {
	return &Controller{
		s: *s,
	}
}

func (ctrl Controller) Get(c *fiber.Ctx) (err error) {
	response := []presenter.Todo{}
	todos, err := ctrl.s.Get()
	if err != nil {
		log.Println(err.Error())

		c.Status(http.StatusInternalServerError)
		return
	}

	copier.Copy(&response, &todos)

	return c.JSON(response)
}

func (ctrl Controller) Create(c *fiber.Ctx) (err error) {
	request := presenter.Todo{}

	if err = c.BodyParser(&request); err != nil {
		return
	}

	entity := todo.Todo{}
	copier.Copy(&entity, &request)

	if _, err = ctrl.s.Create(entity); err != nil {
		return
	}

	return c.JSON(err)
}
