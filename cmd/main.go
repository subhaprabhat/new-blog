package main

import (
	"go-blog/internal/config"
	"go-blog/internal/models"
	"go-blog/pkg/post/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.ConnectDB()
	config.DB.AutoMigrate(&models.Post{})

	routes.PostRoutes(app)

	app.Listen(":3030")
}
