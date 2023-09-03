package main

import (
	"github.com/cozkul/umai/server/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	// Read: https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	setupRoutes(app)

	app.Listen(":3000")
}
