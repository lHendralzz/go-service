package user

import (
	userRepository "go-service/repository/user"

	log "github.com/sirupsen/logrus"
)

// UserService defines methods to interact with the business logic.
type UserService interface {
	Login(Email string, password string) (string, error)
}

// Option define configuration in userService
type Option struct {
	JWTKey string `env:"JWT_KEY"`
}

// userService is a concrete implementation of UserService.
type userService struct {
	userRepository userRepository.UserRepository
	opt            Option
	logger         *log.Logger
}

// NewUserService returns a new UserService with a repository injected.
func NewUserService(repo userRepository.UserRepository, logger *log.Logger, opt Option) UserService {
	return &userService{
		userRepository: repo,
		opt:            opt,
		logger:         logger,
	}
}
