package models

type Teams struct {
	Model
	UsersId uint             `gorm:""`
	Teams   []TeamsCharacter `json:"teams" gorm:"foreignkey:TeamsID"`
}

type TeamsCharacter struct {
	Model
	TeamsID     uint `json:"-"`
	CharacterId uint `json:"characterId"`
	Parameter
}

type BattleTeams struct {
	Model
	Teams []BattleCharacter `json:"teams" gorm:"foreignkey:BattleTeamsID"`
}

type BattleCharacter struct {
	Model
	BattleTeamsID uint       `json:"-"`
	CharacterId   uint       `json:"-"`
	Character     Characters `json:"character" gorm:"-"`
	Parameter
}
