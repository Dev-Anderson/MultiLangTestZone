package repository

import (
	"fmt"
	"gm/internal/database"
	"gm/pkg/models"
)

func CreateUser(user *models.User) error {
	result := database.DBGorm.Create(user)
	return result.Error
}

func GetAllUser() ([]models.User, error) {
	var users []models.User
	if err := database.DBGorm.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("Falha ao buscar o usuario: %w", err)
	}
	return users, nil
}

func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := database.DBGorm.First(&user, id)
	return &user, result.Error
}

func UpdateUser(user *models.User) error {
	result := database.DBGorm.Save(user)
	return result.Error
}

func DeleteUser(id uint) error {
	result := database.DBGorm.Delete(&models.User{}, id)
	return result.Error
}
