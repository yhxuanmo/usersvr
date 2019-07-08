package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name 	 string  `gorm:"type:varchar(100)"`
	Email 	 string  `gorm:"unique;not null"`
	Password string  `gorm:"not null"`
	Icon 	 string  `gorm:"type:varchar(100)"`
}
