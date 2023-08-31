package handlers

import (
	"github.com/cozkul/umai/database"
	"github.com/cozkul/umai/models"
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Hello")
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	// BodyParser from Go Fiber
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&fact)

	// 200 success status code
	return c.Status(200).JSON(fact)
}

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)
	return c.Status(200).JSON(facts)
}
