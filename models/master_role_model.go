package models

import "gorm.io/gorm"

type MasterRole struct {
	gorm.Model
	RoleCode     string `gorm:"type:varchar(100);not null"`
	RoleName     string `gorm:"type:varchar(100);not null"`
	ActiveStatus bool   `gorm:"default:true"`
}
