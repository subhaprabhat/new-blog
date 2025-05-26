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

func CreatePost(c *fiber.Ctx) error {
	post := new(models.Post)

	if err := c.BodyParser(post); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	validate := validator.New()
	errs := validate.Struct(post)
	if errs != nil {
		errMap := customValidate(errs)
		if len(errMap) != 0 {
			return c.Status(400).JSON(errMap)
		}
	}

	config.DB.Create(&post)
	return c.Status(201).JSON(post)
}

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

func DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Post not found"})
	}
	config.DB.Delete(&post)
	return c.SendStatus(204)
}
