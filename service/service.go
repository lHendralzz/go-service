package service

import (
	"go-service/repository"
	"go-service/service/user"

	log "github.com/sirupsen/logrus"
)

// Usecase usecase struct
type Service struct {
	User user.UserService
}

type Options struct {
	User user.Option
}

func Init(repo *repository.Repository, logger *log.Logger, option Options) *Service {
	return &Service{
		User: user.NewUserService(repo.User, logger, option.User),
	}
}
