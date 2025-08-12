package database

import (
	"users-service/models"

	"gorm.io/gorm"
)

type IUserDB interface {
	Create(user *models.User) (*models.User, error)
}
type UserDb struct {
	DB *gorm.DB
}

func NewUserDB(db *gorm.DB) IUserDB {
	return &UserDb{db}
}

func (udb *UserDb) Create(user *models.User) (*models.User, error) {
	if err := udb.DB.AutoMigrate(&models.User{}); err != nil {
		return nil, err
	}
	tx := udb.DB.Create(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}
