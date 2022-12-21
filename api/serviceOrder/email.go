package serviceOrder

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"gopkg.in/mail.v2"
)

type Email struct {
	Sn          string `json:"sn"`
	Ico         int64  `json:"ico"`
	Brand       string `json:"brand"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Model       string `json:"model"`
	Message     string `json:"message"`
	Address     string `json:"address"`
	CompanyName string `json:"companyName"`
}

func SetupEmail(app *fiber.App) {
	email := fiber.New()

	email.Post("/", sendMail)
	app.Mount("/api/service-order/emails", email)
}

func sendMail(c *fiber.Ctx) error {
	emailData := new(Email)
	if err := c.BodyParser(emailData); err != nil {
		response := fiber.Map{"status": 500, "data": map[string]any{}, "err": map[int]any{10: err.Error()}}
		c.Status(500)
		log.Warn().Msgf("%v", response)
		return c.JSON(response)
	}
	clientSn := emailData.Sn
	if clientSn == "" {
		response := fiber.Map{"status": 400, "data": map[string]any{}, "err": map[int]any{4: "Empty SN"}}
		log.Warn().Msgf("%v", response)
		c.Status(400)
		return c.JSON(response)
	}
	clientBrand := emailData.Brand
	if clientBrand == "" {
		response := fiber.Map{"status": 400, "data": map[string]any{}, "err": map[int]any{3: "Empty brand"}}
		log.Warn().Msgf("%v", response)
		c.Status(400)
		return c.JSON(response)
	}
	clientEmail := emailData.Email
	if clientEmail == "" {
		response := fiber.Map{"status": 400, "data": map[string]any{}, "err": map[int]any{4: "Empty email"}}
		log.Warn().Msgf("%v", response)
		c.Status(400)
		return c.JSON(response)
	}
	clientPhone := emailData.Phone
	if clientPhone == "" {
		response := fiber.Map{"status": 400, "data": map[string]any{}, "err": map[int]any{5: "Empty phone"}}
		log.Warn().Msgf("%v", response)
		c.Status(400)
		return c.JSON(response)
	}
	clientModel := emailData.Model
	if clientModel == "" {
		response := fiber.Map{"status": 400, "data": map[string]any{}, "err": map[int]any{6: "Empty model"}}
		log.Warn().Msgf("%v", response)
		c.Status(400)
		return c.JSON(response)
	}
	clientMessage := emailData.Message
	if clientMessage == "" {
		response := fiber.Map{"status": 400, "data": map[string]any{}, "err": map[int]any{7: "Empty message"}}
		c.Status(400)
		log.Warn().Msgf("%v", response)
		return c.JSON(response)
	}
	clientAddress := emailData.Address
	if clientAddress == "" {
		response := fiber.Map{"status": 400, "data": map[string]any{}, "err": map[int]any{8: "Empty address"}}
		log.Warn().Msgf("%v", response)
		c.Status(400)
		return c.JSON(response)
	}
	clientCompanyName := emailData.CompanyName
	if clientCompanyName == "" {
		response := fiber.Map{"status": 400, "data": map[string]any{}, "err": map[int]any{9: "Empty company name"}}
		c.Status(400)
		log.Warn().Msgf("%v", response)
		return c.JSON(response)
	}
	body := fmt.Sprintf("Jméno firmy: %s\nIČO: %d\nE-mail: %s\nTelefonní číslo: %s\nAdresa: %s\nZnačka stroje: %s\nModel stroje: %s\nVýrobní číslo stroje: %s\nPopis závady stroje: %s", emailData.CompanyName, emailData.Ico, emailData.Email, emailData.Phone, emailData.Address, emailData.Brand, emailData.Model, emailData.Sn, emailData.Message)

	email := mail.NewMessage()
	email.SetHeader("From", "servis@rpktechnik.cz")
	email.SetHeader("To", "servis@rpktechnik.cz")
	email.SetHeader("Subject", "Nová objednávka servisu")
	email.SetBody("text/plain", body)

	host := os.Getenv("MAILGUN_SMTP_SERVER")
	port, _ := strconv.ParseInt(os.Getenv("MAILGUN_SMTP_PORT"), 10, 16)
	user := os.Getenv("MAILGUN_SMTP_LOGIN")
	pass := os.Getenv("MAILGUN_SMTP_PASSWORD")

	smtp := mail.NewDialer(host, int(port), user, pass)
	err := smtp.DialAndSend(email)
	if err != nil {
		response := fiber.Map{"status": 500, "data": map[string]any{}, "err": map[int]any{10: err.Error()}}
		c.Status(500)
		log.Warn().Msgf("%v", response)
		return c.JSON(response)
	}
	response := fiber.Map{"status": 200, "data": map[string]any{"Sn": emailData.Sn, "Ico": emailData.Ico, "Brand": emailData.Brand, "Email": emailData.Email, "Phone": emailData.Phone, "Model": emailData.Model, "Message": emailData.Message, "Address": emailData.Address, "CompanyName": emailData.CompanyName}, "err": map[int]any{}}
	log.Info().Msgf("%v", response)
	c.Status(200)
	return c.JSON(response)
}
