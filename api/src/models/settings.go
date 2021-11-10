package models

import "gorm.io/gorm"

type SettingGroup struct {
	gorm.Model
	GroupName     string             `json:"group_name"`
	GroupCategory string             `json:"group_category"`
	Settings      []SettingGroupItem `json:"settings" gorm:"foreignKey:SettingGroupID"`
}

type SettingGroupItem struct {
	gorm.Model
	SettingGroupID uint   `json:"-"`
	SettingName    string `json:"setting_name"`
	SettingLabel   string `json:"setting_label"`
	SettingValue   string `json:"setting_value"`
	SettingType    string `json:"setting_type" gorm:"type: enum('string', 'int', 'boolean'); default:'string'; not null"`
}
