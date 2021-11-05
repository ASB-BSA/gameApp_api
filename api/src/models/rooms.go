package models

import "gorm.io/gorm"

type Rooms struct {
	gorm.Model
	RoomNumber   int    `json:"room_number"`
	RoomStatus   string `json:"room_status" gorm:"type: enum('open', 'close'); default:'open'; not null"`
	UsersID      uint   `json:"-"`
	Users        *Users `json:"create_user"`
	OpponentId   uint   `json:"-"`
	OpponentUser *Users `json:"opponent_user" gorm:"foreignKey:OpponentId"`
}
