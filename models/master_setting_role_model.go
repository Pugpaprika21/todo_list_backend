package models

import "gorm.io/gorm"

type MasterSettingRole struct {
	gorm.Model
	UserID       uint
	MasterRoleID uint
	RefTable     string     `gorm:"type:varchar(150);not null"`
	RefField     string     `gorm:"type:varchar(150);not null"`
	ActiveStatus bool       `gorm:"default:true"`
	MasterRoles  MasterRole `gorm:"foreignKey:MasterRoleID;references:ID"`
	Users        User       `gorm:"foreignKey:UserID;references:ID"`
}
