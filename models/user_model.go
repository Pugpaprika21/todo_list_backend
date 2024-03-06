package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username      string `gorm:"type:varchar(100);not null"`
	Password      string `gorm:"type:varchar(100);not null"`
	FirstName     string `gorm:"type:varchar(150);not null"`
	LastName      string `gorm:"type:varchar(100);not null"`
	ContactNumber string `gorm:"type:varchar(50);not null"`
	Email         string `gorm:"type:varchar(150);not null"`
	Address       string `gorm:"type:varchar(200);not null"`
	FileSyt       string `gorm:"type:varchar(50);not null"`
	Token         string `gorm:"type:varchar(50);not null"`
}
