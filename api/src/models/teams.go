package models

import "gorm.io/gorm"

type Teams struct {
	gorm.Model
	UserId uint             `gorm:""`
	Teams  []TeamsCharacter `json:"teams" gorm:"foreignkey:TeamId"`
}

type TeamsCharacter struct {
	gorm.Model
	TeamId      uint       `json:"-"`
	CharacterId uint       `json:"-"`
	Character   Characters `json:"character" gorm:"-"`
}
