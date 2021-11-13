package controllers

import "github.com/gofiber/fiber/v2"

func Images(c *fiber.Ctx) error {

	return c.SendFile("./assets/images/"+c.Params("id"), false)
}
