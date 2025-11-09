package repositories

import (
	"brand-collab-tracker/models"
	"brand-collab-tracker/config"
)

type TaskInput struct {
	ProjectID uint `json:"project_id" binding:"required"`
	TaskName string `json:"task_name" binding:"required"`
	IsCompleted bool `json:"is_completed"`
	CompletionDate string `json:"completion_date"`
	Notes string `json:"notes"`
}

func CreateTask(input TaskInput) (*models.Task, error) {
	task := models.Task {
		ProjectID: input.ProjectID,
		TaskName: input.TaskName,
		IsCompleted: input.IsCompleted,
		CompletionDate: input.CompletionDate,
		Notes: input.Notes,
	}

	result := config.DB.Create(&task)
	if result.Error != nil {
		return nil, result.Error
	}

	config.DB.Preload("Project").First(&task, task.ID)
	return &task, nil
}

func GetTaskByProject(projectID uint) ([]models.Task, error) {
	var tasks []models.Task
	err := config.DB.Where("project_id = ?", projectID).Preload("Project").Find(&tasks).Error
	return tasks, err
}

func GetTaskByID(id uint) (*models.Task, error) {
	var task models.Task
	err := config.DB.Preload("Project").First(&task, id).Error
	return &task, err
}

func UpdateTask(id uint, input TaskInput) (*models.Task, error) {
	task, err := GetTaskByID(id)
	if err != nil {
		return nil, err
	}

	task.TaskName = input.TaskName
	task.IsCompleted = input.IsCompleted
	task.CompletionDate = input.CompletionDate
	task.Notes = input.Notes

	result := config.DB.Save(task)
	if result.Error != nil {
		return nil, result.Error
	}

	config.DB.Preload("Project").First(&task, task.ID)
	return task, nil
}

func DeleteTask(id uint) error {
	return config.DB.Delete(&models.Task{}, id).Error
}