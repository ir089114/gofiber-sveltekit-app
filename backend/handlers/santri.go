package handlers

import (
	"backend/database"
	"backend/models"

	"github.com/gofiber/fiber/v2"
)

// Get all Santris
func GetSantris(c *fiber.Ctx) error {
	var santris []models.Santri
	database.DB.Find(&santris)
	return c.JSON(santris)
}

// Create new Santri
func CreateSantri(c *fiber.Ctx) error {
	santri := new(models.Santri)
	if err := c.BodyParser(santri); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	database.DB.Create(&santri)
	return c.JSON(santri)
}
