package main

import (
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	err := migrate()
	if err != nil {
		panic(err.Error())
	}
	start()
}
