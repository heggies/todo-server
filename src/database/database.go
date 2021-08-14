package database

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func GetInstance() (*gorm.DB, error) {
	if db == nil {
		var err error
		config := &gorm.Config{}
		if os.Getenv("ENV") == "development" {
			config = &gorm.Config{
				Logger: logger.Default.LogMode(logger.Info),
			}
		}
		db, err = gorm.Open(sqlite.Open("todo.db"), config)
		if err != nil {
			return db, err
		}
	}

	return db, nil
}
