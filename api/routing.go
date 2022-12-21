package api

import (
	"HerokuFirstExperiments/api/career"
	"HerokuFirstExperiments/api/contact"
	"HerokuFirstExperiments/api/serviceOrder"
	"HerokuFirstExperiments/api/test"
	"github.com/gofiber/fiber/v2"
)

func SetupRouting(app *fiber.App) {
	serviceOrder.SetupEmail(app)
	contact.SetupEmail(app)
	career.SetupEmail(app)
	test.SetupTests(app)
	setupCalc(app)
}
