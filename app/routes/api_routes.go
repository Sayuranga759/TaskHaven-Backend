package routes

import (
	"github.com/Sayuranga759/TaskHaven-Backend/app/routes/handler"
	"github.com/gofiber/fiber/v2"
)

func APIRoutes(app *fiber.App) {

	route := app.Group("/todo/v1")

	// health endpoints
	route.Get("livez", handler.Lives)
	route.Get("readyz", handler.Readyz)

	route.Post("/register", handler.UserRegistrationHandler)
}