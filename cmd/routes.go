package main

import (
	"github.com/cozkul/umai/handlers"
	"github.com/gofiber/fiber/v2" // Fiber is an Express.js styled HTTP web framework running on Fasthttp
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)

	app.Post("/fact", handlers.CreateFact)
}
