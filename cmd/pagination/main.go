package main

import (
	"log"

	orders "github.com/maxkondel/go-task-orders"

	"github.com/gofiber/fiber"
)

func main() {
	// Connect with database
	if err := orders.Connect(); err != nil {
		log.Fatal(err)
	}

	// Create a Fiber app
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
