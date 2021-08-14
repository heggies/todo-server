package todo

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/heggies/todo-server/src/controller/todo/presenter"
	"github.com/heggies/todo-server/src/usecase/todo"
	"github.com/heggies/todo-server/src/util/response"
	"github.com/heggies/todo-server/src/util/validator"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type Controller struct {
	s todo.Service
}

func NewController(s *todo.Service) *Controller {
	return &Controller{
		s: *s,
	}
}

func (ctrl *Controller) Get(c *fiber.Ctx) error {
	res := []presenter.Todo{}
	todos, err := ctrl.s.Get()
	if err != nil {
		log.Println(err.Error())
		return response.Error(c, http.StatusInternalServerError)
	}

	copier.Copy(&res, &todos)

	return response.JSON(c, http.StatusOK, res)
}

func (ctrl *Controller) Create(c *fiber.Ctx) error {
	request := presenter.Todo{}

	if err := c.BodyParser(&request); err != nil {
		return response.Error(c, http.StatusUnprocessableEntity)
	}

	errors, err := validator.ValidateStruct(request)
	if err != nil {
		log.Println(err.Error())
		return response.Error(c, http.StatusInternalServerError)
	}

	if len(errors) > 0 {
		return response.Error(c, http.StatusBadRequest, errors...)
	}

	entity, err := ctrl.s.Create(request)
	if err != nil {
		log.Println(err.Error())
		return response.Error(c, http.StatusInternalServerError)
	}

	copier.Copy(&request, &entity)

	return response.JSON(c, http.StatusOK, request)
}

func (ctrl *Controller) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if id <= 0 {
		return response.Error(c, http.StatusBadRequest)
	}

	request := presenter.Todo{}
	if err := c.BodyParser(&request); err != nil {
		log.Println(err.Error())
		return response.Error(c, http.StatusInternalServerError)
	}

	request.ID = id

	entity, err := ctrl.s.Update(request)
	if errors.Is(gorm.ErrRecordNotFound, err) {
		return response.Error(c, http.StatusNotFound)
	} else if err != nil {
		log.Println(err.Error())
		return response.Error(c, http.StatusInternalServerError)
	}

	copier.Copy(&request, entity)

	return response.JSON(c, http.StatusOK, request)
}

func (ctrl *Controller) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if id <= 0 {
		return response.Error(c, http.StatusBadRequest)
	}

	err := ctrl.s.Delete(id)
	if errors.Is(gorm.ErrRecordNotFound, err) {
		return response.Error(c, http.StatusNotFound)
	} else if err != nil {
		log.Println(err.Error())
		return response.Error(c, http.StatusInternalServerError)
	}

	return response.JSON(c, http.StatusOK)
}
