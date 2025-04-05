package middleware

import (
	"nanoshell/database/models"

	"github.com/gofiber/fiber/v2"
)

// RequireAdmin middleware checks if the user is an admin
func RequireAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get admin status from locals (set by auth middleware)
		user := c.Locals("user").(models.User)

		// Check if admin flag exists and is true
		if !user.Admin {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Admin access required",
			})
		}

		return c.Next()
	}
}
