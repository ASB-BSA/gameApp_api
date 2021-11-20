package controllers

import (
	"boomin_game_api/src/database"
	"boomin_game_api/src/middlewares"
	"boomin_game_api/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type BattlePost struct {
	Name string `json:"name"`
}

func PostBattle(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.Status(400)
		return err
	}

	userId, _ := middlewares.GetUserId(c)

	roomId, _ := strconv.Atoi(data["roomId"])

	battle := models.Battle{
		RoomsID: uint(roomId),
	}

	// tx := database.DB.Begin()

	if res := database.DB.First(&battle); res.Error != nil {
		battle.UsersID = userId
	}

	database.DB.Preload("User").Preload("OpponentUser").First(&battle)

	return c.JSON(battle)
}

func GetBattle(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Success",
	})
}
