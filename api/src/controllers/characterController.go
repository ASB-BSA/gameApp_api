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

func GetCharacter(c *fiber.Ctx) error {
	var character []models.Characters
	var ctx = context.Background()

	result, err := database.Cache.Get(ctx, "character").Result()

	if err != nil {
		if result := database.DB.Find(&character); result.Error != nil {
			return result.Error
		}

		bytes, err := json.Marshal(character)

		if err != nil {
			return err
		}

		if err := database.Cache.Set(ctx, "character", bytes, 30*time.Minute).Err(); err != nil {
			return err
		}
	} else {
		json.Unmarshal([]byte(result), &character)
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
