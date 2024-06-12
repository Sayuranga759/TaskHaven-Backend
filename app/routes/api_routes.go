package routes

import (
	"github.com/Sayuranga759/TaskHaven-Backend/app/routes/handler"
	"github.com/Sayuranga759/TaskHaven-Backend/app/routes/middleware"
	"github.com/gofiber/fiber/v2"
)

func APIRoutes(app *fiber.App) {

	route := app.Group("/todo/v1")

	// health endpoints
	route.Get("livez", handler.Lives)
	route.Get("readyz", handler.Readyz)

	// user endpoints
	user := route.Group("/user")
	user.Post("/register", handler.UserRegistrationHandler)
	user.Post("/login", handler.UserLoginHandler)

	task := route.Group("/task", middleware.TokenValidateMiddleware)
	task.Post("/create", handler.CreateTaskHandler)
	task.Post("/update", handler.UpdateTaskHandler)
}