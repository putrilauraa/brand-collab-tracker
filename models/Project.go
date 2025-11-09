package models

import "gorm.io/gorm"

type Project struct {
    gorm.Model
    ID uint `gorm:"primaryKey"`
    
    BrandID uint `gorm:"not null;index:idx_brand_id,unique;references:ID"`
    Brand Brand 
    
    ProjectName string `gorm:"size:255;not null"`
    TargetCompletionDate string
    Status string `gorm:"size:50;default:'On Going'"`
    ContentType string `gorm:"size:100"`
    AgreedFee float64
    
    Tasks []Task `gorm:"foreignKey:ProjectID"`
    Attachments []ProjectAttachment `gorm:"foreignKey:ProjectID"`
}