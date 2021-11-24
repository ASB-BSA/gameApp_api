package models

type Battle struct {
	Model
	RoomsID         uint        `json:"-"`
	UsersID         uint        `json:"userId"`
	UserTeamsID     uint        `json:"-"`
	UserTeams       BattleTeams `json:"userTeams" gorm:"foreignKey:UserTeamsID"`
	User            Users       `json:"createUser" gorm:"foreignKey:UsersID"`
	OpponentID      uint        `json:"-"`
	OpponentTeamsID uint        `json:"-"`
	OpponentTeams   BattleTeams `json:"opponentTeams" gorm:"foreignKey:OpponentTeamsID"`
	OpponentUser    Users       `json:"opponentUser" gorm:"foreignKey:OpponentID"`
	IsActive        bool        `json:"isActive" gorm:"default:true;"`
}
