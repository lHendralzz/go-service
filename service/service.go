package service

import (
	"go-service/repository"
	"go-service/service/user"
	"log"
)

// Usecase usecase struct
type Service struct {
	User              user.UserService
	logger *log.Logger
}

func Init(repo *repository.Repository) *Service {
	return  &Service{
		User: user.NewUserService(repo.User),
	}
}