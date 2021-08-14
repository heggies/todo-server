package todo

import "github.com/heggies/todo-server/src/entity/v1/todo"

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

func (s *Service) Create(entity todo.Todo) (todo.Todo, error) {
	return s.repo.Create(entity)
}

func (s *Service) Delete(id int) (err error) {
	return s.repo.Delete(id)
}
