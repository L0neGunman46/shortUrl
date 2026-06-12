package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
	app.Get("/api/list", routes.ListURL)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	app := fiber.New()
	app.Use(logger.New())
	setRoutes(app)

	log.Fatal(app.Listen(os.Getenv("PORT")))
}
