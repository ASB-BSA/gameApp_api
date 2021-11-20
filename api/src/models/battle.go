package models

import "gorm.io/gorm"

type Battle struct {
	gorm.Model
	RoomsID      uint  `json:"-"`
	UsersID      uint  `json:"userId"`
	UsersTeam    uint  `json:"userTeam"`
	User         Users `json:"createUser" gorm:"foreignKey:UsersID"`
	OpponentId   uint  `json:"opponentId"`
	OpponentUser Users `json:"opponentUser" gorm:"foreignKey:OpponentId"`
	OpponentTeam uint  `json:"opponentTeam"`
	IsActive     bool  `json:"isActive" gorm:"default:true;"`
}
