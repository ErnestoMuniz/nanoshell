package contollers

import (
	"nanoshell/app/dto"
	"nanoshell/database"
	"nanoshell/database/models"
	"nanoshell/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthController struct {
	db *gorm.DB
}

func NewAuthController() *AuthController {
	return &AuthController{db: database.DB}
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	var req dto.LoginDto
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Find user by email
	var user models.User
	if err := c.db.First(&user, "email = ?", req.Email).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid credentials",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get user",
		})
	}

	// Verify password
	if !utils.VerifyPassword(req.Password, user.Password) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	// Generate JWT token using helper function
	token, err := utils.GenerateToken(user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not generate token",
		})
	}

	return ctx.JSON(fiber.Map{
		"token": token,
		"user":  user,
	})
}
