package models

import "gorm.io/gorm"

type Rooms struct {
	gorm.Model
	RoomNumber int    `json:"room_number"`
	RoomStatus string `json:"room_status" gorm:"type: enum('open', 'close'); default:'open'; not null"`
	UsersID    uint   `json:"-"`
	OpponentId uint   `json:"-"`
}
