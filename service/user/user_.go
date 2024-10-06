package user

import (
	"go-service/stdlib/auth"
	x "go-service/stdlib/error"

	"github.com/palantir/stacktrace"
)

func (s *userService) Login(username string, password string) (string, error) {
	isExists, err := s.userRepository.IsExistsUsernameAndPassword(username, password)
	if err != nil {
		return "", stacktrace.Propagate(err, "Select User With Username = %s AND password = %s", username, password)
	}

	if !isExists {
		return "", stacktrace.NewErrorWithCode(x.ErrorFailedLogin, "user not found")
	}

	token := auth.GenerateToken(username, s.opt.JWTSecret)

	return token, nil
}
