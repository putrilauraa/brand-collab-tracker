package repositories

import (
	"brand-collab-tracker/models"
	"brand-collab-tracker/config"
)

type ProjectInput struct {
	BrandID uint `json:"brand_id" binding:"required"`
	ProjectName string `json:"project_name" binding:"required"`
	TargetCompletionDate string `json:"target_completion_date"`
	Status string `json:"status"`
	ContentType string `json:"content_type"`
	AgreedFee float64 `json:"agreed_fee"`
}

func CreateProject(input ProjectInput) (*models.Project, error) {
	project := models.Project{
		BrandID: input.BrandID,
		ProjectName: input.ProjectName,
		TargetCompletionDate: input.TargetCompletionDate,
		Status: input.Status,
		ContentType: input.ContentType,
		AgreedFee: input.AgreedFee,
	}

	result := config.DB.Create(&project)
	if result.Error != nil {
		return nil, result.Error
	}

	config.DB.Preload("Brand.Category").Preload("Tasks").Preload("Attachments").First(&project, project.ID)
	return &project, nil
}

func GetProjects() ([]models.Project, error) {
	var projects []models.Project
	err := config.DB.Preload("Brand.Category").Preload("Tasks").Preload("Attachments").Find(&projects).Error
	return projects, err
}

func GetProjectByID(id uint) (*models.Project, error) {
	var project models.Project
	// err := config.DB.Preload("Brand.Category").Preload("Tasks").Preload("Attachments").First(&project, id).Error
	err := config.DB.First(&project, id).Error
	return &project, err
}

func UpdateProject(id uint, input ProjectInput) (*models.Project, error) {
	project, err := GetProjectByID(id)
	if err != nil {
		return nil, err
	}

	project.BrandID = input.BrandID
	project.ProjectName = input.ProjectName
	project.TargetCompletionDate = input.TargetCompletionDate
	project.Status = input.Status
	project.ContentType = input.ContentType
	project.AgreedFee = input.AgreedFee

	result := config.DB.Save(project)
	if result.Error != nil {
		return nil, result.Error
	}

	config.DB.Preload("Brand.Category").Preload("Tasks").Preload("Attachments").First(&project, project.ID)
	return project, nil
}

func DeleteProject(id uint) error {
	return config.DB.Delete(&models.Project{}, id).Error
}