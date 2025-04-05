package routes

import (
	"nanoshell/app/contollers"
	dto "nanoshell/app/dto"
	"nanoshell/app/middleware"

	m "github.com/Camada8/mandragora"
	"github.com/gofiber/fiber/v2"
)

// SetupAPIRoutes sets up all API routes
func SetupAPIRoutes(app *fiber.App) {
	api := app.Group("/api")
	userController := contollers.NewUserController()

	// Auth routes
	authController := contollers.NewAuthController()
	api.Post("/auth/login", m.WithValidation(m.ValidationConfig{
		Body: &dto.LoginDto{},
	}), authController.Login)

	// User routes
	users := api.Group("/users", middleware.RequireAuth(), middleware.RequireAdmin())
	users.Get("/", userController.GetUsers)
	users.Get("/:id", userController.GetUser)
	users.Post("/", m.WithValidation(m.ValidationConfig{
		Body: &dto.CreateUserDto{},
	}), userController.CreateUser)
	users.Put("/:id", m.WithValidation(m.ValidationConfig{
		Body: &dto.UpdateUserDto{},
	}), userController.UpdateUser)
	users.Delete("/:id", userController.DeleteUser)
}
