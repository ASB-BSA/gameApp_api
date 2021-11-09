package controllers

import (
	"boomin_game_api/src/database"
	"boomin_game_api/src/models"

	"github.com/gofiber/fiber/v2"
)

func CreateSettingGroup(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["group_name"] == "" || data["group_category"] == "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "You're missing a value.",
		})
	}

	setting := models.SettingGroup{
		GroupName:     data["group_name"],
		GroupCategory: data["group_category"],
	}

	if result := database.DB.Create(&setting); result.Error != nil {
		return result.Error
	}

	return c.JSON(&setting)
}

func GetSettingGroup(c *fiber.Ctx) error {
	var settings []models.SettingGroup

	database.DB.Find(&settings)

	return c.JSON(settings)
}

func PutSettingGroup(c *fiber.Ctx) error {
	return c.JSON("Hello")
}

func DeleteSettingGroup(c *fiber.Ctx) error {
	return c.JSON("Hello")
}

func CreateSettingItem(c *fiber.Ctx) error {
	return c.JSON("Hello")
}

func GetSettingItem(c *fiber.Ctx) error {
	return c.JSON("Hello")
}

func PutSettingItem(c *fiber.Ctx) error {
	return c.JSON("Hello")
}

func DeleteSettingItem(c *fiber.Ctx) error {
	return c.JSON("Hello")
}
