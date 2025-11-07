package models

import "gorm.io/gorm"

type Task struct {
    gorm.Model
    ID uint `gorm:"primaryKey"`
    
    ProjectID uint
    Project Project 
    
    TaskName string `gorm:"size:255;not null"` 
    IsCompleted bool `gorm:"default:false"`
    CompletionDate string
    Notes string `gorm:"type:text"`
}