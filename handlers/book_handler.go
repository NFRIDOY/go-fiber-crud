package handlers

import (
	"go-fiber-crud/config"
	"go-fiber-crud/models"

	"github.com/gofiber/fiber/v2"
)

// Create a new Book
func CreateBook(c *fiber.Ctx) error {
	book := new(models.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	config.DB.Create(&book)
	return c.Status(fiber.StatusCreated).JSON(book)
}

// Get all Books
func GetBooks(c *fiber.Ctx) error {
	var books []models.Book
	config.DB.Find(&books)
	return c.Status(fiber.StatusOK).JSON(books)
}

// Get single Book by ID
func GetBookByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book

	if err := config.DB.First(&book, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Book not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(book)
}

// Update Book by ID
func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book

	if err := config.DB.First(&book, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Book not found",
		})
	}

	type UpdateInput struct {
		Title  string  `json:"title"`
		Author string  `json:"author"`
		Price  float64 `json:"price"`
	}

	var input UpdateInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	config.DB.Model(&book).Updates(input)
	return c.Status(fiber.StatusOK).JSON(book)
}

// Delete Book by ID
func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book

	if err := config.DB.First(&book, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Book not found",
		})
	}

	config.DB.Delete(&book)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Book deleted successfully",
	})
}
