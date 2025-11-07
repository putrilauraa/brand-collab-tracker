package models

import "gorm.io/gorm"

type Brand struct {
    gorm.Model
    ID uint `gorm:"primaryKey"`
    Name string `gorm:"size:255;not null"`
    ContactPerson string `gorm:"size:255"`
    ContactEmail string `gorm:"size:255"`
    
    CategoryID uint
    Category CategoryMaster
    
    Projects []Project `gorm:"foreignKey:BrandID"` 
}