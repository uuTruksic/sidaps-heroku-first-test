package api

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

func setupCalc(app *fiber.App) {
	calc := fiber.New()
	calc.Get("/plus", plus)
	calc.Get("/minus", minus)

	calc.Post("/multi", multi)
	calc.Post("/div", div)

	app.Mount("/api/calc", calc)
}

func plus(c *fiber.Ctx) error {
	a, err := strconv.ParseFloat(c.Query("a"), 64)
	if err != nil {
		c.Status(400)
		response := fiber.Map{"status": 400, "data": map[string]any{"a": c.Query("a")}, "err": map[int]string{0: "First input is not a number!"}}
		log.Println(response, map[string]any{"Request_ip": c.IPs()})
		return c.JSON(response)
	}

	b, err := strconv.ParseFloat(c.Query("b"), 64)
	if err != nil {
		response := fiber.Map{"status": 400, "data": map[string]any{"a": c.Query("a"), "b": c.Query("b")}, "err": map[int]string{0: "Second input is not a number!"}}
		c.Status(400)
		log.Println(response, map[string]any{"Request_ip": c.IPs()})
		return c.JSON(response)
	}

	test := strconv.FormatFloat(a+b, 'f', -1, 64)

	c.Status(200)
	response := fiber.Map{"status": 200, "data": map[string]any{"a": c.Query("a"), "b": c.Query("b"), "response": test}, "err": nil}
	log.Println(response, map[string]any{"Request_ip": c.IPs()})
	return c.JSON(response)
}

func minus(c *fiber.Ctx) error {
	a, err := strconv.ParseFloat(c.Query("a"), 64)
	if err != nil {
		c.Status(400)
		response := fiber.Map{"status": 400, "data": map[string]any{"a": c.Query("a")}, "err": map[int]string{0: "First input is not a number!"}}
		log.Println(response, map[string]any{"Request_ip": c.IPs()})
		return c.JSON(response)
	}

	b, err := strconv.ParseFloat(c.Query("b"), 64)
	if err != nil {
		response := fiber.Map{"status": 400, "data": map[string]any{"a": c.Query("a"), "b": c.Query("b")}, "err": map[int]string{0: "Second input is not a number!"}}
		c.Status(400)
		log.Println(response, map[string]any{"Request_ip": c.IPs()})
		return c.JSON(response)
	}

	test := strconv.FormatFloat(a-b, 'f', -1, 64)

	c.Status(200)
	response := fiber.Map{"status": 200, "data": map[string]any{"a": c.Query("a"), "b": c.Query("b"), "response": test}, "err": nil}
	log.Println(response, map[string]any{"Request_ip": c.IPs()})
	return c.JSON(response)
}

func multi(c *fiber.Ctx) error {
	a, err := strconv.ParseFloat(c.FormValue("a"), 64)
	if err != nil {
		c.Status(400)
		response := fiber.Map{"status": 400, "data": map[string]any{"a": c.FormValue("a")}, "err": map[int]string{0: "First input is not a number!"}}
		log.Println(response, map[string]any{"Request_ip": c.IPs()})
		return c.JSON(response)
	}

	b, err := strconv.ParseFloat(c.FormValue("b"), 64)
	if err != nil {
		response := fiber.Map{"status": 400, "data": map[string]any{"a": c.FormValue("a"), "b": c.FormValue("b")}, "err": map[int]string{0: "Second input is not a number!"}}
		c.Status(400)
		log.Println(response, map[string]any{"Request_ip": c.IPs()})
		return c.JSON(response)
	}

	test := strconv.FormatFloat(a*b, 'f', -1, 64)

	c.Status(200)
	response := fiber.Map{"status": 200, "data": map[string]any{"a": c.FormValue("a"), "b": c.FormValue("b"), "response": test}, "err": nil}
	log.Println(response, map[string]any{"Request_ip": c.IPs()})
	return c.JSON(response)

}

func div(c *fiber.Ctx) error {
	a, err := strconv.ParseFloat(c.FormValue("a"), 64)
	if err != nil {
		c.Status(400)
		response := fiber.Map{"status": 400, "data": map[string]any{"a": c.FormValue("a")}, "err": map[int]string{0: "First input is not a number!"}}
		log.Println(response, map[string]any{"Request_ip": c.IPs()})
		return c.JSON(response)
	}

	b, err := strconv.ParseFloat(c.FormValue("b"), 64)
	if err != nil {
		response := fiber.Map{"status": 400, "data": map[string]any{"a": c.FormValue("a"), "b": c.FormValue("b")}, "err": map[int]string{0: "Second input is not a number!"}}
		c.Status(400)
		log.Println(response, map[string]any{"Request_ip": c.IPs()})
		return c.JSON(response)
	}

	if b == 0 {
		response := fiber.Map{"status": 400, "data": map[string]any{"a": c.FormValue("a"), "b": c.FormValue("b")}, "err": map[int]string{1: "Second input is zero!"}}
		c.Status(400)
		log.Println(response, map[string]any{"Request_ip": c.IPs()})
		return c.JSON(response)
	}

	test := strconv.FormatFloat(a/b, 'f', -1, 64)

	c.Status(200)
	response := fiber.Map{"status": 200, "data": map[string]any{"a": c.FormValue("a"), "b": c.FormValue("b"), "response": test}, "err": nil}
	log.Println(response, map[string]any{"Request_ip": c.IPs()})
	return c.JSON(response)
}
