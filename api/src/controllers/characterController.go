package controllers

import (
	"boomin_game_api/src/database"
	"boomin_game_api/src/models"

	"github.com/gofiber/fiber/v2"
)

func GetCharacter(c *fiber.Ctx) error {
	var character []models.Characters

	if result := database.DB.Find(&character); result.Error != nil {
		return result.Error
	}

	return c.JSON(character)
}
