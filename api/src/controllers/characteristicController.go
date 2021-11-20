package controllers

import (
	"boomin_game_api/src/database"
	"boomin_game_api/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetCharacteristic(c *fiber.Ctx) error {
	var Characteristic []models.Characteristic

	database.DB.Find(&Characteristic)

	return c.JSON(Characteristic)
}

func CreateCharacteristic(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	conditionsValue, _ := strconv.Atoi(data["conditionsValue"])
	howMuch, _ := strconv.Atoi(data["howMuch"])

	characteristic := models.Characteristic{
		Name:                 data["name"],
		Timing:               data["timing"],
		ConditionsParameter:  data["conditionsParameter"],
		ConditionsValue:      conditionsValue,
		ConditionsExpression: data["conditionsExpression"],
		ToWhom:               data["toWhom"],
		Parameter:            data["parameter"],
		Happen:               data["happen"],
		HowMuch:              howMuch,
	}

	if result := database.DB.Create(&characteristic); result.Error != nil {
		return result.Error
	}

	return c.JSON(characteristic)
}
