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


// 
type Options struct {
	User user.Option 
}

func Init(repo *repository.Repository, option Options) *Service {
	return  &Service{
		User: user.NewUserService(repo.User, option.User),
	}
}