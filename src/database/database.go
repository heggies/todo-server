package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func GetInstance() (*gorm.DB, error) {
	if db == nil {
		var err error
		config := &gorm.Config{}
		if os.Getenv("ENV") == "development" {
			config.Logger = logger.Default.LogMode(logger.Info)
		}

		dsn := fmt.Sprintf("host=postgres user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Jakarta",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
		)
		db, err = gorm.Open(
			postgres.New(postgres.Config{
				DSN:                  dsn,
				PreferSimpleProtocol: true,
			}),
			config,
		)
		if err != nil {
			return db, err
		}
	}

	return db, nil
}
