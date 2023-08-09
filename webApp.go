package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func Stat(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("status ok")
}

func main() {
	webApp := fiber.New()
	webApp.Get("/status", Stat)
	logrus.Fatal(webApp.Listen(":80")) //localhost
}
