package todo

import (
	v1 "github.com/heggies/todo-server/src/database/v1"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string
	Description *string
	IsDone      bool
}

func Migrate() (err error) {
	db, err := v1.GetInstance()
	if err != nil {
		return
	}

	err = db.AutoMigrate(&Todo{})

	return
}