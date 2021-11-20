package models

import "gorm.io/gorm"

type Characteristic struct {
	gorm.Model
	Name                 string `json:"name"`
	Timing               string `json:"timing" gorm:"type:enum('start', 'attack', 'damage', 'end'); not null"`
	ConditionsParameter  string `json:"conditionsParameter" gorm:"type: enum('hp', 'damage'); default:'hp'; not null"`
	ConditionsValue      int    `json:"conditionsValue"`
	ConditionsExpression string `json:"conditionsExpression" gorm:"type: enum('>', '<', '='); default:'>'; not null"`
	ToWhom               string `json:"toWhom" gorm:"type: enum('myself', 'all_allies', 'random_allies', 'all_enemies', 'ramdom_enemies'); default:'myself'; not null"`
	Parameter            string `json:"parameter" gorm:"type: enum('attack', 'defence', 'critical_rate', 'agility', 'hp', 'mp'); default:'hp'; not null"`
	Happen               string `json:"happen" gorm:"type: enum('+', '-', '='); default:'+'; not null"`
	HowMuch              int    `json:"howMuch"`
}
