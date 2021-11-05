package models

import "gorm.io/gorm"

type Teams struct {
	gorm.Model
	UsersId uint             `gorm:""`
	Teams   []TeamsCharacter `json:"teams" gorm:"foreignkey:TeamsID"`
}

type TeamsCharacter struct {
	gorm.Model
	TeamsID     uint       `json:"-"`
	CharacterId uint       `json:"-"`
	Character   Characters `json:"character" gorm:"-"`
	Parameter
}

type BattleTeams struct {
	gorm.Model
	Teams []BattleCharacter `json:"teams" gorm:"foreignkey:BattleTeamsID"`
}

type BattleCharacter struct {
	gorm.Model
	BattleTeamsID uint       `json:"-"`
	CharacterId   uint       `json:"-"`
	Character     Characters `json:"character" gorm:"-"`
	Parameter
}
