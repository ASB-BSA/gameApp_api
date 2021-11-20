package controllers

import (
	"boomin_game_api/src/database"
	"boomin_game_api/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetCharacter(c *fiber.Ctx) error {
	var character []models.Characters

	if result := database.DB.Find(&character); result.Error != nil {
		return result.Error
	}

	return c.JSON(character)
}

func PutCharacter(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.Status(400)
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	character := models.Characters{
		Name:    data["name"],
		English: data["english"],
		Img:     data["img"],
		Icon:    data["icon"],
	}

	character.ID = uint(id)

	if res := database.DB.Model(&character).Updates(&character); res.Error != nil {
		return res.Error
	}

	return c.JSON(character)
}
