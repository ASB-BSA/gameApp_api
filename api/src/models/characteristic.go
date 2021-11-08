package models

import "gorm.io/gorm"

type Characteristic struct {
	gorm.Model
	CharacterId          uint
	Name                 string `json:"name"`
	ConditionsParameter  string `gorm:"type: enum('hp', 'damage'); default:'hp'; not null"`
	ConditionsValue      int
	ConditionsExpression string `gorm:"type: enum('>', '<', '='); default:'>'; not null"`
	ToWhom               string `gorm:"type: enum('myself', 'all_allies', 'random_allies', 'all_enemies', 'ramdom_enemies'); default:'myself'; not null"`
	Parameter            string `gorm:"type: enum('attack', 'defence', 'critical_rate', 'agility', 'hp', 'mp'); default:'hp'; not null"`
	Happen               string `gorm:"type: enum('+', '-', '='); default:'+'; not null"`
	HowMuch              int
}
