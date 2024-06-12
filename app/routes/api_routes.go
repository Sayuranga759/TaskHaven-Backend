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

	// task endpoints
	task := route.Group("/task", middleware.TokenValidateMiddleware)
	task.Get("", handler.GetTasksByUserIDHandler)
	task.Post("", handler.CreateTaskHandler)
	task.Put("", handler.UpdateTaskHandler)
	task.Delete("", handler.DeleteTaskHandler)
}