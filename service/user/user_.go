package user

import (
	"errors"
	"go-service/stdlib/auth"
)


func (s *userService)Login(username string, password string)(string, error){
    isExists, err := s.userRepository.IsExistsUsernameAndPassword(username, password)
    if err != nil{
        return "", err
    }

    if !isExists {
        return "", errors.New("Wrong Username Password")
    }

    token := auth.GenerateToken(username, s.opt.JWTSecret)

    return token, nil
}