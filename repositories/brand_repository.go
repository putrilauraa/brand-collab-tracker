package repositories

import (
	"brand-collab-tracker/config"
	"brand-collab-tracker/models"
)

type BrandInput struct {
	Name string `json:"name" binding:"required"`
	ContactPerson string `json:"contact_person"`
	ContactEmail string `json:"contact_email"`
	CategoryID uint `json:"category_id" binding:"required"`
}

func CreateBrand(input BrandInput) (*models.Brand, error) {
	brand := models.Brand{
		Name: input.Name,
		ContactPerson: input.ContactPerson,
		ContactEmail: input.ContactEmail,
		CategoryID: input.CategoryID,
	}

	result := config.DB.Create(&brand)
	if result.Error != nil {
		return nil, result.Error
	}

	config.DB.Preload("Category").First(&brand, brand.ID)
	return &brand, nil
}

func GetBrands() ([]models.Brand, error) {
	var brands []models.Brand
	err := config.DB.Preload("Category").Find(&brands).Error
	return brands, err
}

func GetBrandByID(id uint) (*models.Brand, error) {
	var brand models.Brand
	err := config.DB.Preload("Category").First(&brand, id).Error
	return &brand, err
}

func UpdateBrand(id uint, input BrandInput) (*models.Brand, error) {
	brand, err := GetBrandByID(id)
	if err != nil {
		return nil, err
	}

	brand.Name = input.Name
	brand.ContactPerson = input.ContactPerson
	brand.ContactEmail = input.ContactEmail
	brand.CategoryID = input.CategoryID

	result := config.DB.Save(brand)
	if result.Error != nil {
		return nil, result.Error
	}

	config.DB.Preload("Category").First(&brand, brand.ID)
	return brand, nil
}

func DeleteBrand(id uint) error {
	return config.DB.Delete(&models.Brand{}, id).Error
}