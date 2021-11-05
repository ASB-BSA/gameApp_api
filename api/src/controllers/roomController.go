package controllers

import (
	"boomin_game_api/src/database"
	"boomin_game_api/src/middlewares"
	"boomin_game_api/src/models"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetRoom(c *fiber.Ctx) error {
	rand.Seed(time.Now().UnixNano())
	roomId := rand.Intn(1000000)

	userId, _ := middlewares.GetUserId(c)

	isSuccess := true
	var isRoom models.Rooms

	for isSuccess {
		if roomId > 99999 {
			if result := database.DB.Where("room_status = ?", "open").Where("room_number", roomId).First(&isRoom); result.Error != nil {
				room := models.Rooms{
					UsersID:    userId,
					RoomNumber: roomId,
				}
				database.DB.Create(&room)
				isSuccess = false
			}
		}
	}

	return c.JSON(fiber.Map{
		"room_number": roomId,
	})
}

type RoomPost struct {
	RoomNumber int `json:"room_number"`
}

func PostRoom(c *fiber.Ctx) error {
	p := new(RoomPost)

	if err := c.BodyParser(p); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "パラメーターの値が不正です",
		})
	}

	if p.RoomNumber < 99999 && p.RoomNumber > 1000000 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "存在しないルーム番号です",
		})
	}

	var room models.Rooms
	if result := database.DB.Where("room_status = ?", "open").Where("room_number = ?", p.RoomNumber).First(&room); result.Error != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "存在しないルーム番号です",
		})
	}

	userId, _ := middlewares.GetUserId(c)
	if room.UsersID == userId {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "不正アクセスを検知しました",
		})
	}

	room.OpponentId = userId
	room.RoomStatus = "close"
	database.DB.Save(&room)

	return c.JSON(room)
}