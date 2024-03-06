package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	UserIpID     string `gorm:"type:varchar(50);not null"`
	TodoText     string `gorm:"type:varchar(100);not null"`
	TodoStatus   string `gorm:"type:varchar(10);type:enum('pending','inspect','success')"`
	ActiveStatus bool   `gorm:"default:true"`
}
