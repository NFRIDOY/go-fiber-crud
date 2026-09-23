package main

import (
	"log"
	"os"

	"go-fiber-crud/config"
	"go-fiber-crud/models"
	"go-fiber-crud/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// 1. Connect to Database
	config.ConnectDB()

	// 2. Auto-migrate schema
	if err := config.DB.AutoMigrate(&models.Book{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	// 3. Initialize Fiber app
	app := fiber.New()

	// 4. Register Middleware
	app.Use(logger.New())

	// 5. Register Routes
	routes.SetupRoutes(app)

	// 6. Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}
