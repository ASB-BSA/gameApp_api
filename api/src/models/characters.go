package models

import "gorm.io/gorm"

type Characters struct {
	gorm.Model
	Name    string `json:"name"`
	English string `json:"english"`
	Img     string `json:"img"`
}

type Parameter struct {
	Attack    int `json:"attack"`
	Defence   int `json:"defence"`
	Avoidance int `json:"critical_rate"`
	Hit       int `json:"agility"`
	Hp        int `json:"hp"`
	Mp        int `json:"mp"`
}
