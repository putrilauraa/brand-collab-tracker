package models

import "gorm.io/gorm"

type ProjectAttachment struct {
    gorm.Model
    ID uint `gorm:"primaryKey"`
    
    ProjectID uint
    Project Project 
    
    FileName string `gorm:"size:255;not null"`
    FileUrl string `gorm:"type:text;not null"`
    FileType string `gorm:"size:50"`
    Notes string `gorm:"type:text"`
}