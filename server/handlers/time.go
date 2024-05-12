package handlers

import (
	"time"

	"github.com/cozkul/umai/server/database"
	"github.com/gofiber/fiber/v2"
)

func GetSystemTime(c *fiber.Ctx) error {
	var timeStamp time.Time
	database.DB.Raw("SELECT CURRENT_TIMESTAMP;").Scan(&timeStamp)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"time": timeStamp,
	})
}
