package models

import "gorm.io/gorm"

type BattleLogs struct {
	gorm.Model
	Log string `json:"log"`
}
