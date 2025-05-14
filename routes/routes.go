package routes

import (
	"go-blog/controller"

	"github.com/gofiber/fiber/v2"
)

func PostRoutes(app *fiber.App) {
	api := app.Group("/api/posts")
	api.Get("/", controller.GetPosts)
	api.Get("/:id", controller.GetPost)
	api.Post("/", controller.CreatePost)
	 api.Put("/:id", controller.UpdatePost)
	 api.Delete("/:id", controller.DeletePost)
}
