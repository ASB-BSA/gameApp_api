package controllers

import (
	"boomin_game_api/src/database"
	"boomin_game_api/src/middlewares"
	"boomin_game_api/src/models"
	"boomin_game_api/src/utils"
	"math/rand"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetRoom(c *fiber.Ctx) error {
	rand.Seed(time.Now().UnixNano())
	roomId := rand.Intn(1000000)

	userId, _ := middlewares.GetUserId(c)

	if err := utils.CheckTeams(userId); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "チームが設定されていません",
		})
	}

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

func DeleteRoom(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var room models.Rooms
	room.ID = uint(id)

	database.DB.First(&room)

	userId, _ := middlewares.GetUserId(c)
	if room.UsersID == userId {
		database.DB.Delete(&room)
	}

	return c.JSON(fiber.Map{
		"message": "Success.",
	})
}

type RoomPost struct {
	RoomNumber int `json:"roomNumber"`
}

func PostRoom(c *fiber.Ctx) error {
	p := new(RoomPost)

	userId, _ := middlewares.GetUserId(c)

	if err := utils.CheckTeams(userId); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "チームが設定されていません",
		})
	}

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

	if room.UsersID == userId {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "不正アクセスを検知しました",
		})
	}

	room.OpponentId = userId
	room.RoomStatus = "close"

	tx := database.DB.Begin()

	if res := tx.Save(&room); res.Error != nil {
		tx.Rollback()
		return res.Error
	}

	userTeamId, err := CreateBattleTeam(room.UsersID)
	if err != nil {
		tx.Rollback()
		return err
	}

	opponentTeamId, err := CreateBattleTeam(room.OpponentId)
	if err != nil {
		tx.Rollback()
		return err
	}

	battle := models.Battle{
		RoomsID:         room.ID,
		OpponentID:      room.OpponentId,
		OpponentTeamsID: opponentTeamId,
		UsersID:         room.UsersID,
		UserTeamsID:     userTeamId,
	}

	if res := tx.Create(&battle); res.Error != nil {
		tx.Rollback()
		return res.Error
	}

	token, err := middlewares.GenerateBattleToken(battle.ID)

	if err != nil {
		tx.Rollback()
		return err
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

	tx.Commit()

	data := map[string]string{
		"battleID": strconv.Itoa(int(battle.ID)),
	}

	database.PusherClient.Trigger("battle-room", strconv.Itoa(room.RoomNumber), data)

	database.DB.Preload("User").Preload("UserTeams").Preload("UserTeams.Teams").Preload("OpponentUser").Preload("OpponentTeams").Preload("OpponentTeams.Teams").First(&battle)

	return c.JSON(battle)
}

func CreateBattleTeam(id uint) (uint, error) {
	var teams models.BattleTeams

	tx := database.DB.Begin()

	if res := tx.Create(&teams); res.Error != nil {
		tx.Rollback()
		return 0, res.Error
	}

	myTeams := models.Teams{
		UsersId: id,
	}

	if res := tx.Preload("Teams").First(&myTeams); res.Error != nil {
		tx.Rollback()
		return 0, res.Error
	}

	for _, v := range myTeams.Teams {
		chara := models.BattleCharacter{
			BattleTeamsID:    teams.ID,
			Parameter:        v.Parameter,
			CharacterId:      v.CharacterId,
			CharacteristicId: v.CharacteristicId,
		}

		if result := tx.Create(&chara); result.Error != nil {
			tx.Rollback()
			return 0, result.Error
		}
	}

	tx.Commit()

	return teams.ID, nil
}
