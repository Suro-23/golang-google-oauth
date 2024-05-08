package router

import (
	"github.com/Suro-23/api"
	"github.com/gofiber/fiber/v2"
)

func Set(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hellow woooooooorld!")
	})

	app.Get("/oauth", api.GoogleOAuth)
	app.Get("/callback", api.CallBack)
}
