package main

import (
	"go-blog/config"
	"go-blog/models"
	"go-blog/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.ConnectDB()
	config.DB.AutoMigrate(&models.Post{})

	routes.PostRoutes(app)

	app.Listen(":3030")
}
