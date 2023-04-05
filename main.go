package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"net/http"
	"os/exec"
	"strings"
)

type request struct {
	Code string `json:"code"`
}

func runFormatter(inputCode string) ([]byte, error) {
	cmd := exec.Command("rustfmt")
	cmd.Stdin = strings.NewReader(inputCode)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	fmt.Println(out.String())
	return out.Bytes(), nil
}

func base64Decoder(input string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(input)
}

func base64Encoder(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/formatter/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "true",
			"details": "sorobix-formatter",
			"author":  "Hemanth Krishna <@DarthBenro008>",
			"repository": "https://github." +
				"com/sorobix/sorobix-formatter",
		})
	})
	app.Post("/formatter/", func(c *fiber.Ctx) error {
		var data request
		err := c.BodyParser(&data)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"error": "bad request",
			})
		}
		input, err := base64Decoder(data.Code)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"error": "bad base64 input",
			})
		}
		result, err := runFormatter(string(input))
		if err != nil {
			c.Status(http.StatusNotAcceptable)
			return c.JSON(fiber.Map{
				"error": "bad rust code",
				"logs":  err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"formatted_code": base64Encoder(result),
		})
	})
	app.Listen(":3000")
}
