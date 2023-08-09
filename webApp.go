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

func GetTab(c *fiber.Ctx) error {
	return c.JSON(table)
}

var id int

type Request struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
type Response struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	ID    int    `json:"id"`
}

var table = []Response{}

func TaskReq(c *fiber.Ctx) error {
	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	id++
	table = append(table, Response{
		Name:  req.Name,
		Email: req.Email,
		ID:    id,
	})
	return c.Status(fiber.StatusOK).JSON(Response{
		Name:  req.Name,
		Email: req.Email,
		ID:    id,
	})
}

func main() {
	webApp := fiber.New()
	webApp.Get("/", Stat)
	webApp.Get("/incr", Increase)
	webApp.Get("/decr", Decrease)
	webApp.Get("/tab", GetTab)

	webApp.Post("/task", TaskReq)

	logrus.Fatal(webApp.Listen(":80")) //localhost
}
