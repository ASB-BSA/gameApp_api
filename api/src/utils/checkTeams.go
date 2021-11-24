package utils

import (
	"boomin_game_api/src/database"
	"boomin_game_api/src/models"
	"errors"
	"fmt"
)

func CheckTeams(id uint) error {
	var teams models.Teams

	if res := database.DB.Preload("Teams").Where("users_id = ?", id).First(&teams); res.Error != nil {
		return res.Error
	}

	for _, v := range teams.Teams {
		fmt.Println(v)
		if v.CharacterId == 0 {
			return errors.New("チームが設定されていません。")
		}
	}

	return nil
}
