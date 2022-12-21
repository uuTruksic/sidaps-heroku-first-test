package contact

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"gopkg.in/mail.v2"
)

type Email struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Message       string `json:"message"`
	CompanyName   string `json:"companyName"`
	Service       string `json:"service"`
	MachineNumber string `json:"machineNumber"`
}

func SetupEmail(app *fiber.App) {
	email := fiber.New()

	email.Post("/", sendMail)
	app.Mount("/api/contact/emails", email)
}

func sendMail(c *fiber.Ctx) error {
	emailData := new(Email)
	if err := c.BodyParser(emailData); err != nil {
		response := fiber.Map{"status": 500, "data": map[string]any{}, "err": map[int]any{7: err.Error()}}
		c.Status(500)
		log.Warn().Msgf("%v", response)
		return c.JSON(response)
	}
	clientName := emailData.Name
	if clientName == "" {
		response := fiber.Map{"status": 400, "data": map[string]any{}, "err": map[int]any{1: "Empty name"}}
		log.Warn().Msgf("%v", response)
		c.Status(400)
		return c.JSON(response)
	}
	clientEmail := emailData.Email
	if clientEmail == "" {
		response := fiber.Map{"status": 400, "data": map[string]any{}, "err": map[int]any{2: "Empty email"}}
		log.Warn().Msgf("%v", response)
		c.Status(400)
		return c.JSON(response)
	}
	clientPhone := emailData.Phone
	if clientPhone == "" {
		response := fiber.Map{"status": 400, "data": map[string]any{}, "err": map[int]any{3: "Empty phone"}}
		log.Warn().Msgf("%v", response)
		c.Status(400)
		return c.JSON(response)
	}
	clientService := emailData.Service
	if clientService == "" {
		response := fiber.Map{"status": 400, "data": map[string]any{}, "err": map[int]any{4: "Empty service"}}
		c.Status(400)
		log.Warn().Msgf("%v", response)
		return c.JSON(response)
	}
	clientMessage := emailData.Message
	if clientMessage == "" {
		response := fiber.Map{"status": 400, "data": map[string]any{}, "err": map[int]any{5: "Empty message"}}
		c.Status(400)
		log.Warn().Msgf("%v", response)
		return c.JSON(response)
	}
	clientCompanyName := emailData.CompanyName
	if clientCompanyName == "" {
		response := fiber.Map{"status": 400, "data": map[string]any{}, "err": map[int]any{6: "Empty company name"}}
		c.Status(400)
		log.Warn().Msgf("%v", response)
		return c.JSON(response)
	}
	clientMachineNumber := emailData.MachineNumber

	var body string
	if clientMachineNumber != "" {
		body = fmt.Sprintf("Jméno firmy: %s\nJméno a příjmení: %s\nE-mail: %s\nTelefonní číslo: %s\nTéma kontaktu: %s\nZpráva: %s\nČíslo stroje: %s\n", clientCompanyName, clientName, clientEmail, clientPhone, clientService, clientMessage, clientMachineNumber)
	} else {
		body = fmt.Sprintf("Jméno firmy: %s\nJméno a příjmení: %s\nE-mail: %s\nTelefonní číslo: %s\nTéma kontaktu: %s\nZpráva: %s\n ", clientCompanyName, clientName, clientEmail, clientPhone, clientService, clientMessage)
	}

	email := mail.NewMessage()
	email.SetHeader("From", "servis@rpktechnik.cz")
	email.SetHeader("To", "servis@rpktechnik.cz") //servis@rpktechnik.cz
	email.SetHeader("Subject", "Nová zpráva z kontaktní stránky")
	email.SetBody("text/plain", body)

	host := os.Getenv("MAILGUN_SMTP_SERVER")
	port, _ := strconv.ParseInt(os.Getenv("MAILGUN_SMTP_PORT"), 10, 16)
	user := os.Getenv("MAILGUN_SMTP_LOGIN")
	pass := os.Getenv("MAILGUN_SMTP_PASSWORD")

	smtp := mail.NewDialer(host, int(port), user, pass)
	err := smtp.DialAndSend(email)
	if err != nil {
		response := fiber.Map{"status": 500, "data": map[string]any{}, "err": map[int]any{7: err.Error()}}
		c.Status(500)
		log.Warn().Msgf("%v", response)
		return c.JSON(response)
	}
	response := fiber.Map{"status": 200, "data": map[string]any{"Name": clientName, "Email": clientEmail, "Phone": clientPhone, "Service": clientService, "Message": clientMessage, "CompanyName": clientCompanyName}, "err": map[int]any{}}
	log.Info().Msgf("%v", response)
	c.Status(200)
	return c.JSON(response)
}
