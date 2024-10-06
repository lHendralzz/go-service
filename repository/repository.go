package repository

import (
	"go-service/repository/user"
)

type Repository struct{
	User user.UserRepository
}


func Init() *Repository{
	return &Repository{
		User: user.NewUserRepository(),
	}
}