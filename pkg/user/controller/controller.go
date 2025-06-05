package controller

import (
	"go-blog/internal/config"
	"go-blog/internal/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// func validToken(t *jwt.Token, id string) bool {
// 	n, err := strconv.Atoi(id)
// 	if err != nil {
// 		return false
// 	}

// 	claims := t.Claims.(jwt.MapClaims)
// 	uid := int(claims["user_id"].(float64))

// 	return uid == n
// }

// func validUser(id string, p string) bool {
// 	db := config.DB
// 	var user models.User
// 	db.First(&user, id)
// 	if user.Username == "" {
// 		return false
// 	}
// 	if !CheckPasswordHash(p, user.Password) {
// 		return false
// 	}
// 	return true
// }

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := config.DB
	var user models.User
	db.Find(&user, id)
	if user.Username == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "User found", "data": user})
}

func CreateUser(c *fiber.Ctx) error {
	type NewUser struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	db := config.DB
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})
	}

	user.Password = hash
	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}

	newUser := NewUser{
		Email:    user.Email,
		Username: user.Username,
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": newUser})
}

func UpdateUser(c *fiber.Ctx) error {
	type UpdateUserInput struct {
		Names string `json:"names"`
	}
	var uui UpdateUserInput
	if err := c.BodyParser(&uui); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	id := c.Params("id")

	db := config.DB
	var user models.User
	db.First(&user, id)

	user.Username = uui.Names
	db.Save(&user)

	return c.JSON(fiber.Map{"status": "success", "message": "User successfully updated", "data": user})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := config.DB
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}
	if err := db.Delete(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't delete user", "data": err})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "User successfully deleted", "data": nil})
}
