package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func InitApp() {
	engine := html.New("./web/views", ".html")

	config := fiber.Config{
		Views: engine,
	}

	app := fiber.New(config)
	app.Static("/", "./web/views")

	app.Listen(":")
}
