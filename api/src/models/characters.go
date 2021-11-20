package models

import "gorm.io/gorm"

type Characters struct {
	gorm.Model
	Name    string `json:"name"`
	English string `json:"english"`
	Img     string `json:"img"`
	Icon    string `json:"icon"`
}

type Parameter struct {
	Attack       int `json:"attack"`
	Defence      int `json:"defence"`
	Avoidance    int `json:"avoidance"`
	CriticalRate int `json:"criticalRate"`
	Agility      int `json:"agility"`
	Hp           int `json:"hp"`
	Mp           int `json:"mp"`
}
