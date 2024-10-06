package repository

import (
	"go-service/repository/user"

	"gorm.io/gorm"
)

type Repository struct {
	User user.UserRepository
}

func Init(db *gorm.DB) *Repository {
	return &Repository{
		User: user.NewUserRepository(db),
	}
}
