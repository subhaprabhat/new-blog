package controller

import (
	"go-blog/internal/config"
	"go-blog/internal/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Post not found"})
	}
	if err := c.BodyParser(&post); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	validate := validator.New()
	errs := validate.Struct(&post)
	if errs != nil {
		errMap := customValidate(errs)
		if len(errMap) != 0 {
			return c.Status(400).JSON(errMap)
		}
	}

	config.DB.Save(&post)
	return c.JSON(post)
}
