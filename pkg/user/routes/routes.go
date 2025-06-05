package routes

import (
	"go-blog/pkg/user/controller"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	user := app.Group("/api/user")
	user.Get("/:id", controller.GetUser)
	user.Post("/", controller.CreateUser)
	user.Patch("/:id", controller.UpdateUser)
	user.Delete("/:id", controller.DeleteUser)
}
