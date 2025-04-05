package main

import (
	"nanoshell/database"
	"nanoshell/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Warnf("Failed to load environment variables: %v", err)
	}

	// Initialize database
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Run migrations
	if err := database.AutoMigrate(); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Create new Fiber app
	app := fiber.New(fiber.Config{
		AppName: "NanoShell API",
	})

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())

	// Register routes
	registerRoutes(app)

	// Start server
	log.Fatal(app.Listen(":3000"))
}

func registerRoutes(app *fiber.App) {
	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// Register API routes
	routes.SetupAPIRoutes(app)
}
