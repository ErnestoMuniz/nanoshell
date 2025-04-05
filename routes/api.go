package routes

import (
	"nanoshell/app/contollers"
	dto "nanoshell/app/dto"

	m "github.com/Camada8/mandragora"
	"github.com/gofiber/fiber/v2"
)

// SetupAPIRoutes sets up all API routes
func SetupAPIRoutes(app *fiber.App) {
	api := app.Group("/api")
	userController := contollers.NewUserController()

	// User routes
	api.Get("/users", userController.GetUsers)
	api.Get("/users/:id", userController.GetUser)
	api.Post("/users", m.WithValidation(m.ValidationConfig{
		Body: &dto.CreateUserDto{},
	}), userController.CreateUser)
	api.Put("/users/:id", m.WithValidation(m.ValidationConfig{
		Body: &dto.UpdateUserDto{},
	}), userController.UpdateUser)
	api.Delete("/users/:id", userController.DeleteUser)

	// Auth routes
	authController := contollers.NewAuthController()
	api.Post("/auth/login", m.WithValidation(m.ValidationConfig{
		Body: &dto.LoginDto{},
	}), authController.Login)
}
