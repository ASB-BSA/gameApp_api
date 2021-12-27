package models

type BattleLogs struct {
	Model
	BattleID          uint `json:"-"`
	BattleCharacterID uint `json:"battleCharacterID"`
	// BattleCharacter   BattleCharacter `json:"battleCharacter" gorm:"foreignkey:BattleCharacterID"`
	AttackerID uint `json:"attackerID"`
	// Attacker          BattleCharacter `json:"attacker" gorm:"foreignkey:AttackerID"`
	Parameter      string `json:"parameter" gorm:"type: enum('hp')"`
	LogType        string `json:"logType" gorm:"type: enum('attack')"`
	NumericalValue int    `json:"numerical_value"`
}
