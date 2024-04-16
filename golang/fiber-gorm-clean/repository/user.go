package repository

import (
	"fiber-gorm/database"
	"fiber-gorm/models"
)

func CreateUser(user *models.User) error {
	result := database.DB.Create(user)
	return result.Error
}

func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := database.DB.First(&user, id)
	return &user, result.Error
}

func UpdateUser(user *models.User) error {
	result := database.DB.Save(user)
	return result.Error
}

func DeleteUser(id uint) error {
	result := database.DB.Delete(&models.User{}, id)
	return result.Error
}
