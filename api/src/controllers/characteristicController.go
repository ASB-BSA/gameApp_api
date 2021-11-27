package controllers

import (
	"boomin_game_api/src/database"
	"boomin_game_api/src/models"
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetCharacteristic(c *fiber.Ctx) error {
	var Characteristic []models.Characteristic
	var ctx = context.Background()

	result, err := database.Cache.Get(ctx, "characteristic").Result()

	if err != nil {
		database.DB.Find(&Characteristic)

		bytes, err := json.Marshal(Characteristic)

		if err != nil {
			return err
		}

		if err := database.Cache.Set(ctx, "characteristic", bytes, 30*time.Minute).Err(); err != nil {
			return err
		}
	} else {
		json.Unmarshal([]byte(result), &Characteristic)
	}

	return c.JSON(Characteristic)
}

func CreateCharacteristic(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	conditionsValue, err := strconv.Atoi(data["conditionsValue"])
	if err != nil {
		return err
	}

	howMuch, err := strconv.Atoi(data["howMuch"])
	if err != nil {
		return err
	}

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
