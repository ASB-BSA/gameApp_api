package controllers

import (
	"boomin_game_api/src/database"
	"boomin_game_api/src/middlewares"
	"boomin_game_api/src/models"
	"fmt"
	"sort"
	"time"

	"github.com/gofiber/fiber/v2"
)

type BattlePost struct {
	BattleId int `json:"battleId"`
}

func PostBattle(c *fiber.Ctx) error {
	p := new(BattlePost)

	userId, _ := middlewares.GetUserId(c)

	var battle models.Battle

	battle.ID = uint(p.BattleId)

	if result := database.DB.Where("is_active = ?", 1).First(&battle); result.Error != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "対戦情報が見つかりませんでした",
		})
	}

	if battle.UsersID != userId && battle.OpponentID != userId {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "対戦情報が見つかりませんでした",
		})
	}

	token, err := middlewares.GenerateBattleToken(uint(p.BattleId))

	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "対戦情報が見つかりませんでした",
		})
	}

	cookie := fiber.Cookie{
		Name:     "battle-jwt",
		Value:    token,
		Expires:  time.Now().AddDate(10, 0, 0),
		SameSite: "None",
		Secure:   true,
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	database.DB.Preload("User").Preload("UserTeams").Preload("UserTeams.Teams").Preload("OpponentUser").Preload("OpponentTeams").Preload("OpponentTeams.Teams").First(&battle)

	return c.JSON(battle)
}

func GetBattle(c *fiber.Ctx) error {
	id, _ := middlewares.GetBattleId(c)

	var battle models.Battle
	battle.ID = id

	if result := database.DB.Preload("User").Preload("UserTeams").Preload("UserTeams.Teams").Preload("OpponentUser").Preload("OpponentTeams").Preload("OpponentTeams.Teams").First(&battle); result.Error != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "対戦情報が見つかりませんでした",
		})
	}

	return c.JSON(battle)
}

func CreateBattleLog(c *fiber.Ctx) error {
	id, _ := middlewares.GetBattleId(c)
	var battle models.Battle
	battle.ID = id

	if result := database.DB.Preload("User").Preload("UserTeams").Preload("UserTeams.Teams").Preload("UserTeams.Teams.Characteristics").Preload("OpponentUser").Preload("OpponentTeams").Preload("OpponentTeams.Teams").Preload("OpponentTeams.Teams.Characteristics").Where("is_active = ?", "1").First(&battle); result.Error != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "対戦情報が見つかりませんでした",
		})
	}

	// キャラクターのみを抽出
	characters := battle.OpponentTeams.Teams
	characters2 := battle.UserTeams.Teams
	characters = append(characters, characters2...)

	sort.Slice(characters, func(i, j int) bool {
		return characters[i].Agility > characters[j].Agility
	})

	// スキル発動
	for _, v := range characters {
		if v.Characteristics.Timing == "start" {
			fmt.Println(v.Characteristics)
		}
	}

	return c.JSON(characters)
}
