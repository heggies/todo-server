package todo

import "github.com/heggies/todo-server/src/entity/v1/todo"

type Writer interface {
	Get() (entities []todo.Todo, err error)
}

type Reader interface {
	Create(entity todo.Todo) (todo.Todo, error)
	Update(entity todo.Todo) (todo.Todo, error)
	Delete(id int) (err error)
}

type Repositorier interface {
	Writer
	Reader
}
