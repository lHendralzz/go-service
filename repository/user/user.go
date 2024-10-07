package user

import (
	"go-service/model"

	"gorm.io/gorm"
)

// UserRepository is an interface that defines the data operations.
type UserRepository interface {
	GetPasswordByUsername(username string) (model.User, error)
}

// userRepository is a concrete implementation of the UserRepository interface.
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository returns an instance of the concrete implementation of UserRepository.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
