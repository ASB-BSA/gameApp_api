package controllers

import (
	"boomin_game_api/src/database"
	"boomin_game_api/src/models"
	"strconv"

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

	database.DB.Preload("Settings").Find(&settings)

	return c.JSON(settings)
}

func PutSettingGroup(c *fiber.Ctx) error {
	return c.JSON("Hello")
}

func DeleteSettingGroup(c *fiber.Ctx) error {
	return c.JSON("Hello")
}

func CreateSettingItem(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.Status(400)
		return err
	}

	if data["setting_name"] == "" || data["setting_label"] == "" || data["setting_value"] == "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "You're missing a value.",
		})
	}

	id, _ := strconv.Atoi(c.Params("id"))

	var res models.SettingGroupItem
	if result := database.DB.Where("setting_label = ?", data["setting_label"]).First(&res); result.Error == nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Already registered.",
		})
	}

	setting := models.SettingGroupItem{
		SettingGroupID: uint(id),
		SettingName:    data["setting_name"],
		SettingLabel:   data["setting_label"],
		SettingValue:   data["setting_value"],
		SettingType:    data["setting_type"],
	}

	if result := database.DB.Create(&setting); result.Error != nil {
		return result.Error
	}

	return c.JSON(&setting)
}

func PutSettingItem(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	setting := models.SettingGroupItem{
		SettingName:  data["setting_name"],
		SettingLabel: data["setting_label"],
		SettingValue: data["setting_value"],
		SettingType:  data["setting_type"],
	}

	setting.ID = uint(id)

	database.DB.Model(&setting).Updates(&setting)

	return c.JSON(setting)
}

func DeleteSettingItem(c *fiber.Ctx) error {
	return c.JSON("Hello")
}
