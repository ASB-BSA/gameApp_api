package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name string `json:"name" gorm:"unique_index"`
}
