package controllers

import (
	"boomin_game_api/src/database"
	"boomin_game_api/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetBattlelog(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var battleLog []models.BattleLogs
	var battle models.Battle
	battle.ID = uint(id)

	if result := database.DB.Preload("User").Preload("UserTeams").Preload("UserTeams.Teams").Preload("UserTeams.Teams.Characteristics").Preload("OpponentUser").Preload("OpponentTeams").Preload("OpponentTeams.Teams").Preload("OpponentTeams.Teams.Characteristics").Where("is_active = ?", "1").First(&battle); result.Error != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "対戦情報が見つかりませんでした",
		})
	}

	if res := database.DB.Where("battle_id = ?", id).Find(&battleLog); res.Error != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "対戦情報が見つかりませんでした",
		})
	}

	return c.JSON(fiber.Map{
		"logs":       battleLog,
		"battleData": battle,
	})
}
