package user

import (
	"go-service/model"

	log "github.com/sirupsen/logrus"

	"gorm.io/gorm"
)

// UserRepository is an interface that defines the data operations.
type UserRepository interface {
	GetUserByEmail(email string) (model.User, error)
}

// userRepository is a concrete implementation of the UserRepository interface.
type userRepository struct {
	db     *gorm.DB
	logger *log.Logger
}

// NewUserRepository returns an instance of the concrete implementation of UserRepository.
func NewUserRepository(db *gorm.DB, logger *log.Logger) UserRepository {
	return &userRepository{
		db:     db,
		logger: logger,
	}
}
