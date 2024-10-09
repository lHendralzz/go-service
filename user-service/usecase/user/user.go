package user

import (
	userRepository "go-service/repository/user"

	log "github.com/sirupsen/logrus"
)

// defines methods to interact with the business logic.
type UserUsecase interface {
	Login(Email string, password string) (string, error)
}

// Option define configuration in userService
type Option struct {
	JWTKey string `env:"JWT_KEY"`
}

// a concrete implementation of UserService.
type userUsecase struct {
	userRepository userRepository.UserRepository
	opt            Option
	logger         *log.Logger
}

// NewUserUsecase returns a new UserUsecase with a repository injected.
func NewUserUsecase(repo userRepository.UserRepository, logger *log.Logger, opt Option) UserUsecase {
	return &userUsecase{
		userRepository: repo,
		opt:            opt,
		logger:         logger,
	}
}
