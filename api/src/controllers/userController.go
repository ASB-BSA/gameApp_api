package controllers

import (
	"boomin_game_api/src/database"
	"boomin_game_api/src/middlewares"
	"boomin_game_api/src/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserPost struct {
	Name string `name:"name"`
}

func PostUser(c *fiber.Ctx) error {
	p := new(UserPost)

	if err := c.BodyParser(p); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "パラメーターの値が不正です",
		})
	}

	user := models.Users{
		Name: p.Name,
	}

	if result := database.DB.Create(&user); result.Error != nil {
		return result.Error
	}

	token, err := middlewares.GenerateJWT(user.ID)

	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "登録に失敗しました。",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().AddDate(10, 0, 0),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

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
	database.DB.Where("id = ?", id).First(&user)

	return c.JSON(user)
}
