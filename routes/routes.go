package routes

import (
	"go-fiber-crud/docs"
	"go-fiber-crud/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	docs.Register(app)

	api := app.Group("/api/v1")

	// Book endpoints
	books := api.Group("/books")
	books.Post("/", handlers.CreateBook)
	books.Get("/", handlers.GetBooks)
	books.Get("/:id", handlers.GetBookByID)
	books.Put("/:id", handlers.UpdateBook)
	books.Delete("/:id", handlers.DeleteBook)
}
