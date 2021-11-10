package controllers

import (
	"boomin_game_api/src/database"
	"boomin_game_api/src/middlewares"
	"boomin_game_api/src/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func AdminRegister(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "password do not match",
		})
	}

	var admin models.Admin
	database.DB.Where("user = ?", data["user"]).First(&admin)

	if admin.ID != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "This user name cannot be registered.",
		})
	}

	admin.User = data["user"]

	admin.SetPassword(data["password"])

	database.DB.Create(&admin)

	return c.JSON(admin)
}

func AdminLogin(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.Admin

	database.DB.Where("user = ?", data["user"]).First(&user)

	if user.ID == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	token, err := middlewares.GenerateAdminJWT(user.ID)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	cookie := fiber.Cookie{
		Name:     "admin-jwt",
		Value:    token,
		Expires:  time.Now().AddDate(10, 0, 0),
		SameSite: "None",
		Secure:   true,
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func Admins(c *fiber.Ctx) error {
	var admin []models.Admin

	database.DB.Find(&admin)

	return c.JSON(admin)
}
