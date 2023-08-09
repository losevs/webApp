package main

import (
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

var mutex = &sync.Mutex{}
var count int

func Stat(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("status ok")
}

func Increase(c *fiber.Ctx) error {
	mutex.Lock()
	count++
	mutex.Unlock()
	return c.SendString(strconv.Itoa(count))
}

func Decrease(c *fiber.Ctx) error {
	mutex.Lock()
	count--
	mutex.Unlock()
	return c.SendString(strconv.Itoa(count))
}

func main() {
	webApp := fiber.New()
	webApp.Get("/", Stat)
	webApp.Get("/incr", Increase)
	webApp.Get("/decr", Decrease)
	logrus.Fatal(webApp.Listen(":80")) //localhost
}
