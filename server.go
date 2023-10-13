package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet/v2"
)

func main() {
	// Create a new engine
	engine := jet.New("./views", ".jet")

	// Or from an embedded system
	// See github.com/gofiber/embed for examples
	// engine := jet.NewFileSystem(http.Dir("./views", ".jet"))

	// Pass the engine to the views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	app.Get("/layout", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/main")
	})

	log.Fatal(app.Listen(":3000"))
}
