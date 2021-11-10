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

	conditionsValue, _ := strconv.Atoi(data["conditions_value"])
	howMuch, _ := strconv.Atoi(data["how_much"])

	characteristic := models.Characteristic{
		Name:                 data["name"],
		ConditionsParameter:  data["conditions_parameter"],
		ConditionsValue:      conditionsValue,
		ConditionsExpression: data["conditions_expression"],
		ToWhom:               data["to_whom"],
		Parameter:            data["parameter"],
		Happen:               data["happen"],
		HowMuch:              howMuch,
	}

	if result := database.DB.Create(&characteristic); result.Error != nil {
		return result.Error
	}

	return c.JSON(characteristic)
}
