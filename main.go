package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		version := os.Getenv("APP_VERSION")
		msg := fmt.Sprintf("APP VERSION: %s", version)
		return c.SendString(msg)
	})

	log.Fatal(app.Listen(":3000"))
}
