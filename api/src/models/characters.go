package models

import "gorm.io/gorm"

type Characters struct {
	gorm.Model
	Name    string `json:"name"`
	English string `json:"english"`
	Img     string `json:"img"`
}
