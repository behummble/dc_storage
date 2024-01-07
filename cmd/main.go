package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	InitApp()
}

func InitApp() {
	engine := html.New("./web/views", ".html")

	config := fiber.Config{
		Views: engine,
	}

	app := fiber.New(config)
	app.Static("/", "./web/views")
	app.Static("/styles/", "./web/styles")
	app.Static("/files/", "./web/files")

	app.Listen(":3000")
}
