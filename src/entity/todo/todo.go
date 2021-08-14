package todo

import (
	"github.com/heggies/todo-server/src/database"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string
	Description *string
	IsDone      *bool
}

func Migrate() (err error) {
	db, err := database.GetInstance()
	if err != nil {
		return
	}

	err = db.AutoMigrate(&Todo{})

	return
}
