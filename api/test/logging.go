package test

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"os"
)

func setupLoggingTest(app *fiber.App) {
	logging := fiber.New()

	logging.Get("/", testLogging)

	app.Mount("/api/test/log", logging)
}

func testLogging(c *fiber.Ctx) error {
	log := zerolog.New(os.Stdout)
	log.Warn().Msg("Testovaci log")
	return c.SendString("nothing to see here")
}
