package career

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"gopkg.in/mail.v2"
)

type Email struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Message    string `json:"message"`
	CareerName string `json:"careerName"`
}

func SetupEmail(app *fiber.App) {
	email := fiber.New()

	email.Post("/", sendMail)
	app.Mount("/api/career/emails", email)
}

func sendMail(c *fiber.Ctx) error {
	emailData := Email{
		Name:       c.FormValue("name"),
		Email:      c.FormValue("email"),
		Phone:      c.FormValue("phone"),
		Message:    c.FormValue("message"),
		CareerName: c.FormValue("careerName"),
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
	clientMessage := emailData.Message
	if clientMessage == "" {
		response := fiber.Map{"status": 400, "data": map[string]any{}, "err": map[int]any{5: "Empty message"}}
		c.Status(400)
		log.Warn().Msgf("%v", response)
		return c.JSON(response)
	}
	clientFileHandler, err := c.FormFile("file")
	if err != nil {
		response := fiber.Map{"status": 500, "data": map[string]any{}, "err": map[int]any{6: err.Error()}}
		c.Status(500)
		log.Warn().Msgf("%v", response)
		return c.JSON(response)
	}
	clientCareerName := emailData.CareerName
	if clientCareerName == "" {
		response := fiber.Map{"status": 400, "data": map[string]any{}, "err": map[int]any{7: "Empty Career name"}}
		c.Status(400)
		log.Warn().Msgf("%v", response)
		return c.JSON(response)
	}

	clientFile, err := clientFileHandler.Open()
	if err != nil {
		response := fiber.Map{"status": 500, "data": map[string]any{}, "err": map[int]any{6: err.Error()}}
		c.Status(500)
		log.Warn().Msgf("%v", response)
		return c.JSON(response)
	}

	body := fmt.Sprintf("Název pracovní pozice: %s\nJméno a příjmení: %s\nE-mail: %s\nTelefonní číslo: %s\nZpráva: %s\n", clientCareerName, clientName, clientEmail, clientPhone, clientMessage)

	email := mail.NewMessage()
	email.SetHeader("From", "servis@rpktechnik.cz") //Rpktechnik@sandbox5b70b2e4d4f04712b66f6d03f791c294.mailgun.org
	email.SetHeader("To", "servis@rpktechnik.cz")
	email.SetHeader("Subject", fmt.Sprintf("Nová odpověď na inzerát práce %s", clientCareerName))
	email.SetBody("text/plain", body)
	email.AttachReader(clientFileHandler.Filename, clientFile)

	host := os.Getenv("MAILGUN_SMTP_SERVER")
	port, _ := strconv.ParseInt(os.Getenv("MAILGUN_SMTP_PORT"), 10, 16)
	user := os.Getenv("MAILGUN_SMTP_LOGIN")
	pass := os.Getenv("MAILGUN_SMTP_PASSWORD")

	smtp := mail.NewDialer(host, int(port), user, pass)
	if err := smtp.DialAndSend(email); err != nil {
		response := fiber.Map{"status": 500, "data": map[string]any{}, "err": map[int]any{6: err.Error()}}
		c.Status(500)
		log.Warn().Msgf("%v", response)
		return c.JSON(response)
	}
	response := fiber.Map{"status": 200, "data": map[string]any{"Name": clientName, "Email": clientEmail, "Phone": clientPhone, "Message": clientMessage, "FileName": clientFileHandler.Filename, "CareerName": clientCareerName}, "err": map[int]any{}}
	log.Info().Msgf("%v", response)
	c.Status(200)
	return c.JSON(response)
}
