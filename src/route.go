package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	controller "github.com/heggies/todo-server/src/controller"
)

func start() {
	app := fiber.New()
	app.Use(cors.New())

	if os.Getenv("ENV") == "development" {
		app.Use(logger.New())
	}

	if err := controller.StartHandler(app); err != nil {
		log.Panicln(err.Error())
	}

	app.Listen(":" + os.Getenv("PORT"))
}
