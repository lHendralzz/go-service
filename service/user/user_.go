package user

import (
	"fmt"
	"go-service/stdlib/auth"
	x "go-service/stdlib/error"

	"github.com/palantir/stacktrace"
)

func (s *userService) Login(username string, password string) (string, error) {
	isExists, err := s.userRepository.IsExistsUsernameAndPassword(username, password)
	if err != nil {
		return "", stacktrace.PropagateWithCode(err, x.ErrorQuery, "Select User With Username = %s AND password = %s", username, password)
	}

	if !isExists {
		return "", stacktrace.NewErrorWithCode(x.ErrorLogin, "user not found")
	}

	token, err := auth.GenerateToken(username, s.opt.JWTKey)
	if err != nil {
		fmt.Println(err)
		return "", stacktrace.PropagateWithCode(err, x.ErrorGenerateJWTToken, "failed generateToken for username %s", username)
	}

	return token, nil
}
