package controllers

import (
	"boomin_game_api/src/database"
	"boomin_game_api/src/middlewares"
	"boomin_game_api/src/models"
	"fmt"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetRoom(c *fiber.Ctx) error {
	rand.Seed(time.Now().UnixNano())
	roomId := rand.Intn(1000000)

	userId, _ := middlewares.GetUserId(c)

	isSuccess := true

	var room models.Rooms

	for isSuccess {
		if roomId > 99999 {
			room.RoomStatus = "open"
			room.RoomNumber = roomId

			var count int64

			if res := database.DB.Model(&room).Where("room_states = ?", "open").Count(&count); res.Error != nil && count == 0 {
				isSuccess = false
			} else {
				roomId = rand.Intn(1000000)
			}

			fmt.Println(count)
		} else {
			roomId = rand.Intn(1000000)
		}
	}

	room.UsersID = userId

	if res := database.DB.Create(&room); res.Error != nil {
		return res.Error
	}

	return c.JSON(room)
}

type RoomPost struct {
	RoomNumber int `json:"roomNumber"`
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

	fmt.Println(p.RoomNumber)

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
