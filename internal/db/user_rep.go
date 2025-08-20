package db

import "AuthBase/internal/models"

func CreateUser(r *models.User) error {
	result := DB.Create(&r)
	return result.Error
}

func CheckIfExists(email string) int64 {
	var count int64
	DB.Model(&models.User{}).Where("email = ?", email).Count(&count)
	return count
}
