package todo

import "github.com/heggies/todo-server/src/entity/todo"

type Reader interface {
	Get() (entities []todo.Todo, err error)
}

type Writer interface {
	Create(entity todo.Todo) (todo.Todo, error)
	Update(entity todo.Todo) (todo.Todo, error)
	Delete(id int) (err error)
}

type Repositorier interface {
	Reader
	Writer
}
