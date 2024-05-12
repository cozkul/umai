package main

import (
	"github.com/cozkul/umai/server/config"
	"github.com/cozkul/umai/server/database"
	"github.com/cozkul/umai/server/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		return
	}
	database.ConnectDb()
	database.RunMigrations()

	helpers.InitializeUniverse()

	app := fiber.New()
	// Read: https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	setupRoutes(app)

	app.Listen(":3000")

	helpers.SetUpServerTick()
}
