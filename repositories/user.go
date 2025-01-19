package repositories

import (
	"go-gin-tasks-manager/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (repo *UserRepository) GetById(id string) (*models.User, error) {
	return nil, nil
}
