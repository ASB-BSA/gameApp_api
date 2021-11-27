package models

type Rooms struct {
	Model
	RoomNumber int    `json:"roomNumber"`
	RoomStatus string `json:"roomStatus" gorm:"type: enum('open', 'close'); default:'open'; not null"`
	UsersID    uint   `json:"-"`
	OpponentId uint   `json:"-"`
}
