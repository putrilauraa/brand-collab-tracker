package repositories

import (
	"brand-collab-tracker/config"
	"brand-collab-tracker/models"
)

type AttachmentInput struct {
	ProjectID uint `json:"project_id" binding:"required"`
	FileName string `json:"file_name" binding:"required"`
	FileUrl string `json:"file_url" binding:"required"`
	FileType string `json:"file_type"`
	Notes string `json:"notes"`
}

func CreateAttachment(input AttachmentInput) (*models.ProjectAttachment, error) {
	attachment := models.ProjectAttachment{
		ProjectID: input.ProjectID,
		FileName: input.FileName,
		FileUrl: input.FileUrl,
		FileType: input.FileType,
		Notes: input.Notes,
	}

	result := config.DB.Create(&attachment)
	if result.Error != nil {
		return nil, result.Error
	}
	
	config.DB.Preload("Project").First(&attachment, attachment.ID)
	return &attachment, nil
}

func GetAttachmentsByProject(projectID uint) ([]models.ProjectAttachment, error) {
	var attachments []models.ProjectAttachment
	err := config.DB.Where("project_id = ?", projectID).Preload("Project").Find(&attachments).Error
	return attachments, err
}

func GetAttachmentByID(id uint) (*models.ProjectAttachment, error) {
	var attachment models.ProjectAttachment
	err := config.DB.Preload("Project").First(&attachment, id).Error
	return &attachment, err
}

func UpdateAttachment(id uint, input AttachmentInput) (*models.ProjectAttachment, error) {
	attachment, err := GetAttachmentByID(id)
	if err != nil {
		return nil, err
	}

	attachment.FileName = input.FileName
	attachment.FileUrl = input.FileUrl
	attachment.FileType = input.FileType
	attachment.Notes = input.Notes

	result := config.DB.Save(attachment)
	if result.Error != nil {
		return nil, result.Error
	}
	
	config.DB.Preload("Project").First(&attachment, attachment.ID)
	return attachment, nil
}

func DeleteAttachment(id uint) error {
	return config.DB.Delete(&models.ProjectAttachment{}, id).Error
}