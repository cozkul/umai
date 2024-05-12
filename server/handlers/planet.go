package handlers

import (
	"errors"
	"strconv"

	"github.com/cozkul/umai/server/database"
	"github.com/cozkul/umai/server/models"
	"github.com/gofiber/fiber/v2"
)

func getPlanet(user *models.User, planetString string, preload string) (*models.Planet, error) {
	u64, err := strconv.ParseUint(planetString, 10, 32)
	if err != nil {
		return nil, errors.New("error finding planet")
	}
	planetID := uint(u64)

	var planet models.Planet
	result := database.DB.Preload(preload).Where("id = ?", planetID).First(&planet)
	if result.Error != nil {
		return nil, errors.New("error finding planet")
	}
	if planet.UserID == nil || *planet.UserID != user.ID {
		return nil, errors.New("planet does not belong to user")
	}
	return &planet, nil
}

func GetPlanetBuildings(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	planet, err := getPlanet(user, c.Params("planet"), "Buildings")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(planet.Buildings)
}

func BuildingUpgrade(c *fiber.Ctx) error {
	return nil
}

func BuildingDemolish(c *fiber.Ctx) error {
	return nil
}

func GetPlanetResources(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	planet, err := getPlanet(user, c.Params("planet"), "Resource")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(planet.Resource)
}
