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
		sslMode := "require"
		if os.Getenv("ENV") == "development" {
			config.Logger = logger.Default.LogMode(logger.Info)
			sslMode = "disable"
		}

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=%s TimeZone=Asia/Jakarta",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
			sslMode,
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
