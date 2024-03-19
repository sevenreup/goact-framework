package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sevenreup/goact"
	"html/template"
	"log"
)

func main() {
	opts := goact.GoactEngineOpts{
		OutputDir:        "./dist",
		WorkingDir:       "./views",
		IsDebug:          true,
		StructPath:       "./dto",
		TsTypeOutputPath: "./views/types",
		DefaultRenderData: goact.RenderData{
			Title: "My App",
			Head:  template.HTML(` <link rel="stylesheet" href="/main.css">`),
		},
		InjectCss: false,
	}
	engine := goact.CreateGoactEngine(&opts)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("main/index.tsx", goact.RenderData{
			Title: "Hello, World!",
			Props: map[string]string{
				"title": "Hello, World!",
			},
		})
	})

	err := app.Listen(":3000")

	if err != nil {
		log.Panic(err)
	}
}
