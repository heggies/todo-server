package main

import "github.com/heggies/todo-server/src/entity/v1/todo"

func migrate() (err error) {
	err = todo.Migrate()
	return
}
