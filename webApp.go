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
	return c.JSON(ppl)
}

type Request struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
type Response struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	ID    int    `json:"id"`
}

var id int
var ppl = map[int]Response{}

func TaskReq(c *fiber.Ctx) error {
	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	id++
	hum := Response{
		Name:  req.Name,
		Email: req.Email,
		ID:    id,
	}
	ppl[id] = hum
	return c.Status(fiber.StatusOK).JSON(hum)
}

func PatchReq(c *fiber.Ctx) error {
	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	needId, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if _, isHere := ppl[int(needId)]; !isHere {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	changedTask := Response{
		Name:  req.Name,
		Email: req.Email,
		ID:    int(needId),
	}
	ppl[int(needId)] = changedTask
	return c.Status(fiber.StatusOK).JSON(changedTask)
}

func DelReq(c *fiber.Ctx) error {
	needId, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	delete(ppl, int(needId))
	return c.SendStatus(fiber.StatusOK)
}

func main() {
	webApp := fiber.New()
	webApp.Get("/", Stat)
	webApp.Get("/incr", Increase)
	webApp.Get("/decr", Decrease)
	webApp.Get("/tab", GetTab)

	webApp.Post("/task", TaskReq)

	webApp.Patch("/patch/:id", PatchReq)

	webApp.Delete("/del/:id", DelReq)

	logrus.Fatal(webApp.Listen(":80")) //localhost
}
