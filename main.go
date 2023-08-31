package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go-wedding/db"
	"go-wedding/routers"
)

func main() {

	var err error

	DB := db.Init()

	app := fiber.New()

	app.Get("/metrics", monitor.New())

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(requestid.New())

	routers.SetupRoutes(app, DB)

	err = app.Listen(":3000")

	if err != nil {
		panic(err)
	}
}
