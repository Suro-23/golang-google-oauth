package main

import (
	"github.com/Suro-23/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	router.Set(app)
	app.Listen(":8080")
}
