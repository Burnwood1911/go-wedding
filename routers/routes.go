package routers

import (
	"github.com/gofiber/fiber/v2"
	"go-wedding/handlers"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api/v1/invite")

	api.Post("/create", handlers.NewHandler(db).CreateInvite)

	api.Post("/scan", handlers.NewHandler(db).ScanInvite)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

}
