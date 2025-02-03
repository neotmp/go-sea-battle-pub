package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/neotmp/go-sea-battle/api"
)

func Setup(app *fiber.App) {

	app.Get("/api/check", api.Check)
	app.Get("/api/start", api.Start)
	app.Post("/api/play", api.Play)
	app.Get("/api/resume", api.Resume)
}
