package user

import (
	model "go-service/model"
	userRepository "go-service/repository/user"
)

// UserService defines methods to interact with the business logic.
type UserService interface {
    GetUserProfile(id int) (*model.User, error)
}

// userService is a concrete implementation of UserService.
type userService struct {
    UserRepository userRepository.UserRepository
}

// NewUserService returns a new UserService with a repository injected.
func NewUserService(repo userRepository.UserRepository) UserService {
    return &userService{
        UserRepository: repo,
    }
}