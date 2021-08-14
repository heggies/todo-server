package todo

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/heggies/todo-server/src/controller/v1/todo/presenter"
	"github.com/heggies/todo-server/src/entity/v1/todo"
	usecase "github.com/heggies/todo-server/src/usecase/v1/todo"
	"github.com/heggies/todo-server/src/util/response"
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
	res := []presenter.Todo{}
	todos, err := ctrl.s.Get()
	if err != nil {
		log.Println(err.Error())

		c.Status(http.StatusInternalServerError)
		return
	}

	copier.Copy(&res, &todos)

	return response.JSON(c, res)
}

func (ctrl Controller) Create(c *fiber.Ctx) (err error) {
	request := presenter.Todo{}

	if err = c.BodyParser(&request); err != nil {
		return
	}

	entity := todo.Todo{}
	copier.Copy(&entity, &request)

	entity, err = ctrl.s.Create(entity)
	if err != nil {
		return
	}

	copier.Copy(&request, &entity)

	return response.JSON(c, request)
}
