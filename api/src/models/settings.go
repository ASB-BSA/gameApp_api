package models

import "gorm.io/gorm"

type SettingGroup struct {
	gorm.Model
	Settings      []Setting `json:"settings"`
	GroupName     string    `json:"group_name"`
	GroupCategory string    `json:"group_category"`
}

type Setting struct {
	gorm.Model
	SettingGroupID uint   `json:"group_id"`
	SettingName    string `json:"setting_name"`
	SettingLabel   string `json:"setting_label"`
	SettingValue   string `json:"setting_value"`
	SettingType    string `json:"setting_type" gorm:"type: enum('string', 'int', 'boolean'); default:'string'; not null"`
}
