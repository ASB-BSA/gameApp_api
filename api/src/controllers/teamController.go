package controllers

import (
	"boomin_game_api/src/database"
	"boomin_game_api/src/middlewares"
	"boomin_game_api/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func PutTeamCharacter(c *fiber.Ctx) error {
	var data map[string]int

	if err := c.BodyParser(&data); err != nil {
		c.Status(400)
		return err
	}
	userId, _ := middlewares.GetUserId(c)

	var team models.Teams
	team.ID = uint(data["teamId"])
	team.UsersId = userId

	if res := database.DB.First(&team); res.Error != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "存在しないIDです",
		})
	}

	id, _ := strconv.Atoi(c.Params("id"))

	chara := models.TeamsCharacter{
		CharacterId: uint(data["charaId"]),
	}

	chara.ID = uint(id)
	chara.CharacterId = uint(data["characterId"])
	chara.Attack = data["attack"]
	chara.Defence = data["defence"]
	chara.Avoidance = data["avoidance"]
	chara.CriticalRate = data["criticalRate"]
	chara.Agility = data["agility"]
	chara.Hp = data["hp"]
	chara.Mp = data["mp"]

	if res := database.DB.Model(&chara).Updates(&chara); res.Error != nil {
		return res.Error
	}

	return c.JSON(chara)
}
