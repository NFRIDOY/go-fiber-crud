package docs

import (
	_ "embed"

	"github.com/gofiber/fiber/v2"
)

//go:embed openapi.json
var specification []byte

//go:embed index.html
var swaggerUI []byte

// Register serves the embedded API documentation alongside the API.
func Register(app *fiber.App) {
	app.Get("/swagger", func(c *fiber.Ctx) error {
		return c.Redirect("/swagger/index.html", fiber.StatusTemporaryRedirect)
	})
	app.Get("/swagger/index.html", func(c *fiber.Ctx) error {
		return c.Type("html").Send(swaggerUI)
	})
	app.Get("/swagger/openapi.json", func(c *fiber.Ctx) error {
		return c.Type("json").Send(specification)
	})
}
