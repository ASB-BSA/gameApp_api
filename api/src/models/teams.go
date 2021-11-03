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
	Attack      int        `json:"attack"`
	Defence     int        `json:"defence"`
	Avoidance   int        `json:"critical_rate"`
	Hit         int        `json:"agility"`
	Hp          int        `json:"hp"`
	Mp          int        `json:"mp"`
}
