package v1

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetInstance() (*gorm.DB, error) {
	if db == nil {
		var err error
		db, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
		if err != nil {
			return db, err
		}
	}

	return db, nil
}
