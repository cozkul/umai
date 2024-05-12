package handlers

import (
	"github.com/cozkul/umai/server/database"
	"github.com/cozkul/umai/server/models"
	"github.com/gofiber/fiber/v2"
)

func GetUniverseSystem(c *fiber.Ctx) error {
	systemID := c.Params("system")
	var system []models.System
	result := database.DB.Where("system_id = ?", systemID).Find(&system)
	if result.Error != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "Unable to get system",
			"error":   result.Error.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(system)
}
