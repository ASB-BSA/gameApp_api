package models

type Characters struct {
	Model
	Name    string `json:"name"`
	English string `json:"english"`
	Img     string `json:"img"`
	Icon    string `json:"icon"`
}

type Parameter struct {
	Attack       int `json:"attack" gorm:"default:50"`
	Defence      int `json:"defence" gorm:"default:50"`
	Avoidance    int `json:"avoidance" gorm:"default:50"`
	CriticalRate int `json:"criticalRate" gorm:"default:50"`
	Agility      int `json:"agility" gorm:"default:50"`
	Hp           int `json:"hp" gorm:"default:200"`
	Mp           int `json:"mp" gorm:"default:0"`
}
