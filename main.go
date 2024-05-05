package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Failed to load .env")
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Payment subscription engine"})
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "3001"
	}

	app.Listen(fmt.Sprintf(":%v", port))
}
