package main

import (
	"HerokuFirstExperiments/api"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
)

func main() {
	app := fiber.New(fiber.Config{ErrorHandler: fiberErrorHandling})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	app.Static("/", "./build")
	api.SetupRouting(app)

	port := os.Getenv("PORT")
	app.Listen(fmt.Sprintf(":%s", port))
}

func fiberErrorHandling(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	if code == fiber.StatusNotFound {
		return c.SendFile("./build/index.html")
	}
	return nil
}
