package models

import "gorm.io/gorm"

type CategoryMaster struct {
    gorm.Model
    ID uint `gorm:"primaryKey"`
    Name string `gorm:"size:100;unique;not null"`
}