package models

import "gorm.io/gorm"

type Characters struct {
	gorm.Model
	Name string `json:"name"`
	Img  string `json:"img"`
}
