package middleware

import (
	"nanoshell/database"
	"nanoshell/database/models"
	"nanoshell/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// RequireAuth is a middleware that checks for a valid JWT token
func RequireAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header required",
			})
		}

		// Check if the header starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid authorization format",
			})
		}

		// Extract the token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Decode and validate the token
		claims, err := utils.DecodeToken(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		// Get user from database
		var user models.User
		if err := database.DB.First(&user, "id = ?", claims["user_id"]).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "User not found",
			})
		}

		// Check if user is active
		if !user.Active {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "User is inactive",
			})
		}

		// Store user and claims in context
		c.Locals("user", user)

		return c.Next()
	}
}
