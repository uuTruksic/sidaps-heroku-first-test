package test

import (
	"github.com/gofiber/fiber/v2"
)

func SetupTests(app *fiber.App) {
	setupLoggingTest(app)
	setupEmail(app)
}
