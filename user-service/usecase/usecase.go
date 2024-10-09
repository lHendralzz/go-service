package usecase

import (
	"go-service/repository"
	"go-service/stdlib/redis"
	"go-service/usecase/user"

	log "github.com/sirupsen/logrus"
)

// Usecase usecase struct
type Usecase struct {
	User user.UserUsecase
}

type Options struct {
	User user.Option
}

func Init(repo *repository.Repository, logger *log.Logger, option Options, redis *redis.RedisLock) *Usecase {
	return &Usecase{
		User: user.NewUserUsecase(repo.User, logger, option.User),
	}
}
