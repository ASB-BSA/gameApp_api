package models

type BattleLogs struct {
	Model
	BattleID          uint   `json:"-"`
	BattleCharacterID uint   `json:"characterID"`
	AttackerID        uint   `json:"attackerID"`
	Parameter         string `json:"parameter" gorm:"type: enum('hp')"`
	LogType           string `json:"logType" gorm:"type: enum('attack')"`
	NumericalValue    int    `json:"numerical_value"`
}
