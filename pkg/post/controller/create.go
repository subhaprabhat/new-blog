package controller

import (
	"go-blog/internal/config"
	"go-blog/internal/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

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

	err := config.DB.Create(&post).Error
	if err != nil {
		return err
	}

	return c.Status(201).JSON(post)
}
