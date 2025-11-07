package repositories

import (
	"brand-collab-tracker/config"
	"brand-collab-tracker/models"
)

func CreateCategory(category *models.CategoryMaster) error {
	return config.DB.Create(category).Error
}

func GetCategoryByID(id uint) (models.CategoryMaster, error) {
	var category models.CategoryMaster
	err := config.DB.First(&category, id).Error
	return category, err
}

func GetCategories() ([]models.CategoryMaster, error) {
	var categories []models.CategoryMaster
	err := config.DB.Find(&categories).Error
	return categories, err
}

func UpdateCategory(category *models.CategoryMaster) error {
	return config.DB.Save(category).Error
}

func DeleteCategory(id uint) error {
	return config.DB.Delete(&models.CategoryMaster{}, id).Error
}