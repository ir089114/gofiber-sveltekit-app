package routes

import (
	"backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Santri routes
	api.Get("/santris", handlers.GetSantris)
	api.Post("/santris", handlers.CreateSantri)
}
