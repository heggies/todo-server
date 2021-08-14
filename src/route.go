package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	v1 "github.com/heggies/todo-server/src/controller/v1"
)

func start() {
	app := fiber.New()
	app.Use(cors.New())

	if os.Getenv("ENV") == "development" {
		app.Use(logger.New())
	}

	if err := v1.StartHandler(app); err != nil {
		log.Panicln(err.Error())
	}

	app.Listen(":3000")
}