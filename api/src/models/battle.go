package models

import "gorm.io/gorm"

type Battle struct {
	gorm.Model
	UsersID      uint   `json:"-"`
	UsersTeam    uint   `json:"user_team"`
	Users        *Users `json:"create_user"`
	OpponentId   uint   `json:"-"`
	OpponentUser *Users `json:"opponent_user" gorm:"foreignKey:OpponentId"`
	OpponentTeam uint   `json:"opponent_team"`
	IsActive     bool   `json:"is_active"`
}
