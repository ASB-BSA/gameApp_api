package models

type BattleLogs struct {
	Model
	BattleID          uint   `json:"-"`
	BattleCharacterID uint   `json:"-"`
	Parameter         string `json:"parameter"`
	NumericalValue    int    `json:"numerical_value"`
}
