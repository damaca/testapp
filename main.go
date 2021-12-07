package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// GET /api/register
	app.Get("/", func(c *fiber.Ctx) error {
		version := os.Getenv("APP_VERSION")
		msg := fmt.Sprintf("APP VERSION: %s", version)
		return c.SendString(msg) // => ✋ register
	})

	log.Fatal(app.Listen(":3000"))
}
