package todo

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/heggies/todo-server/src/controller/v1/todo/presenter"
	"github.com/heggies/todo-server/src/entity/v1/todo"
	usecase "github.com/heggies/todo-server/src/usecase/v1/todo"
	"github.com/heggies/todo-server/src/util/response"
	"github.com/heggies/todo-server/src/util/validator"
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
		return response.Error(c, http.StatusInternalServerError)
	}

	copier.Copy(&res, &todos)

	return response.JSON(c, res)
}

func (ctrl Controller) Create(c *fiber.Ctx) (err error) {
	request := presenter.Todo{}

	if err = c.BodyParser(&request); err != nil {
		return
	}

	errors, err := validator.ValidateStruct(request)
	if err != nil {
		log.Println(err.Error())
		return response.Error(c, http.StatusInternalServerError)
	}

	if len(errors) > 0 {
		return response.Error(c, http.StatusBadRequest, errors...)
	}

	entity := todo.Todo{
		Title:       request.Title,
		Description: &request.Description,
	}

	entity, err = ctrl.s.Create(entity)
	if err != nil {
		log.Println(err.Error())
		return response.Error(c, http.StatusInternalServerError)
	}

	copier.Copy(&request, &entity)

	return response.JSON(c, request)
}
