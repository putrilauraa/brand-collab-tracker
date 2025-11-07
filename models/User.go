package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    ID uint `gorm:"primaryKey"`
    Username string `gorm:"size:100;unique;not null"`
    Password string `gorm:"size:255;not null"`
}