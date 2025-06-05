package controller

import (
	"fmt"
	"go-blog/internal/config"
	"go-blog/internal/models"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func customValidate(err error) map[string]string {
	errors := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrors {
			tag := strings.ToLower(fieldErr.Tag())
			field := strings.ToLower(fieldErr.Field())

			switch tag {
			case "required":
				errors[field] = fmt.Sprintf("%v is a required field", field)
			default:
				errors[field] = fieldErr.Error()
			}
		}
	}
	return errors
}

func GetPosts(c *fiber.Ctx) error {
	var posts []models.Post
	config.DB.Order("id desc").Find(&posts)
	return c.JSON(posts)
}

func GetPost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Post
	result := config.DB.First(&post, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Post not found"})
	}
	return c.JSON(post)
}

