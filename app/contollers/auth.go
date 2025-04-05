package contollers

import (
	"nanoshell/app/dto"
	"nanoshell/database"
	"nanoshell/database/models"
	"nanoshell/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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

	// Generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["admin"] = user.Admin
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("your-secret-key")) // TODO: Use env variable
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not generate token",
		})
	}

	return ctx.JSON(fiber.Map{
		"token": t,
		"user":  user,
	})

}
