package contollers

import (
	"nanoshell/database"
	"nanoshell/database/models"
	"nanoshell/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserController struct {
	db *gorm.DB
}

func NewUserController() *UserController {
	return &UserController{
		db: database.DB,
	}
}

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=255"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Admin    bool   `json:"admin"`
}

type UpdateUserRequest struct {
	Username string `json:"username" validate:"omitempty,min=3,max=255"`
	Email    string `json:"email" validate:"omitempty,email"`
	Password string `json:"password" validate:"omitempty,min=8"`
	Admin    *bool  `json:"admin"`
	Active   *bool  `json:"active"`
}

func (c *UserController) CreateUser(ctx *fiber.Ctx) error {
	var req CreateUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Hash the password
	hashedPassword := utils.HashPassword(req.Password)

	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		Admin:    req.Admin,
		Active:   true,
	}

	if err := c.db.Create(&user).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(user)
}

func (c *UserController) GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var user models.User

	if err := c.db.First(&user, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get user",
		})
	}

	return ctx.JSON(user)
}

func (c *UserController) GetUsers(ctx *fiber.Ctx) error {
	var users []models.User

	if err := c.db.Find(&users).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get users",
		})
	}

	return ctx.JSON(users)
}

func (c *UserController) UpdateUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var user models.User

	// First, check if user exists
	if err := c.db.First(&user, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get user",
		})
	}

	// Parse update data
	var req UpdateUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Update fields if they are provided
	updates := make(map[string]interface{})

	if req.Username != "" {
		updates["username"] = req.Username
	}
	if req.Email != "" {
		updates["email"] = req.Email
	}
	if req.Password != "" {
		updates["password"] = utils.HashPassword(req.Password)
	}
	if req.Admin != nil {
		updates["admin"] = *req.Admin
	}
	if req.Active != nil {
		updates["active"] = *req.Active
	}

	// Only update if there are changes
	if len(updates) > 0 {
		if err := c.db.Model(&user).Updates(updates).Error; err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update user",
			})
		}
	}

	return ctx.JSON(user)
}

func (c *UserController) DeleteUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var user models.User

	// First, check if user exists
	if err := c.db.First(&user, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get user",
		})
	}

	// Delete user
	if err := c.db.Delete(&user).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete user",
		})
	}

	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
