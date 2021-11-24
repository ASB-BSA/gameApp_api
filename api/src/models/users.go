package models

type Users struct {
	Model
	Name  string `json:"name" gorm:"unique_index"`
	Teams Teams  `gorm:"foreignkey:UsersId"`
}
