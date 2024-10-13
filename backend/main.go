package main

import (
	"backend/database"
	"backend/routes"

	"github.com/gofiber/fiber/v2"
)

func init() {
	// Connect to the database
	database.ConnectDB()
	database.Migration()

}

func main() {
	app := fiber.New()

	// Setup routes
	routes.SetupRoutes(app)

	app.Listen(":3000")
}
