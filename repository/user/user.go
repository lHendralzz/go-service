package user

import (
	model "go-service/model"
)

// UserRepository is an interface that defines the data operations.
type UserRepository interface {
    GetUserByID(id int) (*model.User, error)
}

// userRepository is a concrete implementation of the UserRepository interface.
type userRepository struct {
}

// NewUserRepository returns an instance of the concrete implementation of UserRepository.
func NewUserRepository() UserRepository {
    return &userRepository{}
}
