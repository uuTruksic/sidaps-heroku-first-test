package test

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
)

func setupEmail(app *fiber.App) {
	email := fiber.New()

	email.Post("/", sendMail)
	app.Mount("/api/test/emails", email)
}

func sendMail(c *fiber.Ctx) error {
	clientName := c.FormValue("name")
	clientEmail := c.FormValue("email")
	clientPhone := c.FormValue("phone")
	clientMessage := c.FormValue("message")
	body := fmt.Sprintf("Zpráva od %s (email: %s telefon: %s) poslal přes stránku rpktechnik.cz zprávu\n%s", clientName, clientEmail, clientPhone, clientMessage)

	email := gomail.NewMessage()
	email.SetHeader("From", "alias@alias.alias")
	email.SetHeader("To", "sidaps.project@gmail.com")
	email.SetHeader("Subject", "Testing subject")
	email.SetBody("text/plain", body)

	host := os.Getenv("MAILGUN_SMTP_SERVER")
	port, _ := strconv.ParseInt(os.Getenv("MAILGUN_SMTP_PORT"), 10, 16)
	user := os.Getenv("MAILGUN_SMTP_LOGIN")
	pass := os.Getenv("MAILGUN_SMTP_PASSWORD")

	smtp := gomail.NewDialer(host, int(port), user, pass)
	err := smtp.DialAndSend(email)
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.SendString("email sent via gomail")
}
