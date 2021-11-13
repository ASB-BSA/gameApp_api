package controllers

import (
	"boomin_game_api/src/database"
	"boomin_game_api/src/middlewares"
	"boomin_game_api/src/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserPost struct {
	Name string `json:"name"`
}

func PostUser(c *fiber.Ctx) error {
	p := new(UserPost)

	if err := c.BodyParser(p); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "パラメーターの値が不正です",
		})
	}

	if p.Name == "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "名前を入力してください",
		})
	}

	user := models.Users{
		Name: p.Name,
	}

	tx := database.DB.Begin()

	if result := tx.Create(&user); result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	teams := models.Teams{
		UsersId: user.ID,
	}

	if result := tx.Create(&teams); result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	for i := 0; i < 5; i++ {
		chara := models.TeamsCharacter{
			TeamsID: teams.ID,
		}

		if result := tx.Create(&chara); result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	}

	token, err := middlewares.GenerateJWT(user.ID)

	if err != nil {
		tx.Rollback()
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "登録に失敗しました。",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().AddDate(10, 0, 0),
		SameSite: "None",
		Secure:   true,
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	tx.Commit()

	return c.JSON(&user)
}

func GetUser(c *fiber.Ctx) error {
	id, err := middlewares.GetUserId(c)

	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "ユーザー認証に失敗しました。",
		})
	}

	var user models.Users
	database.DB.Where("id = ?", id).Preload("Teams").Preload("Teams.Teams").First(&user)

	return c.JSON(user)
}
