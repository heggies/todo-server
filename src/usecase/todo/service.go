package todo

import (
	"github.com/heggies/todo-server/src/controller/todo/presenter"
	"github.com/heggies/todo-server/src/entity/todo"
	"gorm.io/gorm"
)

type Service struct {
	repo Repositorier
}

func NewService(repo Repositorier) (service *Service, err error) {
	service = &Service{
		repo: repo,
	}
	return
}

func (s *Service) Get() (entities []todo.Todo, err error) {
	entities, err = s.repo.Get()
	return
}

func (s *Service) Create(entity presenter.Todo) (todo.Todo, error) {
	return s.repo.Create(todo.Todo{
		Title:       entity.Title,
		Description: &entity.Description,
	})
}

func (s *Service) Update(entity presenter.Todo) (todo.Todo, error) {
	return s.repo.Update(todo.Todo{
		Model: gorm.Model{
			ID: uint(entity.ID),
		},
		Title:       entity.Title,
		Description: &entity.Description,
		IsDone:      &entity.IsDone,
	})
}

func (s *Service) Delete(id int) (err error) {
	return s.repo.Delete(id)
}
