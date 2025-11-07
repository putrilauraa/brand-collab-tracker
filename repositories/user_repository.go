package repositories

import (
	"errors"

	"brand-collab-tracker/config"
	"brand-collab-tracker/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserAuthRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func RegisterUser(data UserAuthRequest) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Username: data.Username,
		Password: string(hashedPassword),
	}

	result := config.DB.Create(&user)
	return &user, result.Error
}

func VerifyUser(data UserAuthRequest) (*models.User, error) {
	var user models.User

	result := config.DB.Where("username = ?", data.Username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("username not found")
		}
		return nil, result.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return nil, errors.New("wrong password")
	}

	return &user, nil
}