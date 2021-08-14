package main

import "github.com/heggies/todo-server/src/entity/todo"

func migrate() (err error) {
	err = todo.Migrate()
	return
}
