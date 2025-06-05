package main

import (
	"go-blog/internal/config"
	"go-blog/internal/models"
	postRoutes "go-blog/pkg/post/routes"
	userRoutes "go-blog/pkg/user/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})

	postRoutes.PostRoutes(app)

	userRoutes.UserRoutes(app)

	if err := app.Listen(":3030"); err != nil {
		panic(err)
	}

}
