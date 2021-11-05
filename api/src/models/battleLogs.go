package models

import "gorm.io/gorm"

type BattleLogs struct {
	gorm.Model
	BattleID          uint   `json:"-"`
	BattleCharacterID uint   `json:"-"`
	Parameter         string `json:"parameter"`
	NumericalValue    int    `json:"numerical_value"`
}
