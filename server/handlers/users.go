package handlers

import (
	"errors"

	"github.com/cozkul/umai/server/database"
	"github.com/cozkul/umai/server/models"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&user)

	// Success status code
	return c.Status(200).JSON(user)
}

func Authenticate(c *fiber.Ctx) error {
	return errors.New("No way Jose.")
}
