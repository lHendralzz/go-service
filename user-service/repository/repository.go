package repository

import (
	"go-service/repository/user"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository struct {
	User user.UserRepository
}

func Init(db *gorm.DB, logger *log.Logger) *Repository {
	return &Repository{
		User: user.NewUserRepository(db, logger),
	}
}
