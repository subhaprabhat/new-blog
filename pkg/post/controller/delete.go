package controller

import (
	"go-blog/internal/config"
	"go-blog/internal/models"

	"github.com/gofiber/fiber/v2"
)

func DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Post not found"})
	}
	config.DB.Delete(&post)
	return c.SendStatus(204)
}
