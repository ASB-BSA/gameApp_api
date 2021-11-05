package controllers

import "github.com/gofiber/fiber/v2"

type BattlePost struct {
	Name string `json:"name"`
}

func PostBattle(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Success",
	})
}

func GetBattle(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Success",
	})
}
