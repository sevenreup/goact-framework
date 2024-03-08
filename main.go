package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sevenreup/goact"
	"log"
)

func main() {
	engine := goact.CreateGoatEngine("./dist", "./views")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("./views/main/index.tsx", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	err := app.Listen(":3000")

	if err != nil {
		log.Panic(err)
	}
}
