package services

import (
	"api-instagram/app/models"
	"api-instagram/db"
	"errors"
	"fmt"
)

type UserCriteria string

const (
	ByID       UserCriteria = "id"
	ByUsername UserCriteria = "username"
	ByEmail    UserCriteria = "email"
	ByPhone    UserCriteria = "phone"
)

func FindUserByCriteria(identifier string, criteria UserCriteria) (*models.Users, error) {
	var user models.Users
	var condition string

	switch criteria {
	case ByID:
		condition = "id = ?"
	case ByUsername:
		condition = "username = ?"
	case ByPhone:
		condition = "phone = ?"
	default:
		return nil, errors.New("invalid criteria")
	}

	if err := db.Instance.Table("Users").Where(condition, identifier).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func InsertUser(user *models.Users) (*models.Users, error) {
	if err := db.Instance.Table("Users").Create(&user).Error; err != nil {
		fmt.Print(err)
		return nil, err
	}

	return user, nil
}

func GetUsers() ([]models.Users, error) {
	var users []models.Users

	if err := db.Instance.Table("Users").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
