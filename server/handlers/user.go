package handlers

import (
	"github.com/cozkul/umai/server/database"
	"github.com/cozkul/umai/server/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func GetUserSystems(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	result := database.DB.Preload(clause.Associations).Find(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Unable to get user systems",
			"error":   result.Error.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(user.Systems)
}

func GetUserPlanets(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	result := database.DB.Preload(clause.Associations).Find(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Unable to get user systems",
			"error":   result.Error.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(user.Planets)
}
