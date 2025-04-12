package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/prongbang/microservice-go/libs/common"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		message := common.Hello("Auth Service")
		return c.SendString(message)
	})

	fmt.Println("🚀 Auth Service running on http://localhost:9002")
	log.Fatal(app.Listen(":9002"))
}
