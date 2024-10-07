package user

import (
	"fmt"
	"go-service/stdlib/auth"
	x "go-service/stdlib/error"
	"go-service/stdlib/hash"

	"github.com/palantir/stacktrace"
)

func (s *userService) Login(username string, password string) (string, error) {
	user, err := s.userRepository.GetPasswordByUsername(username)
	if err != nil {
		return "", stacktrace.PropagateWithCode(err, x.ErrorQuery, "Select User With Username = %s AND password = %s", username, password)
	}

	isValid := hash.ComparePassword(user.Password, password)
	if !isValid {
		return "", stacktrace.NewErrorWithCode(x.ErrorLogin, "username or password is invalid")
	}

	token, err := auth.GenerateToken(username, s.opt.JWTKey)
	if err != nil {
		fmt.Println(err)
		return "", stacktrace.PropagateWithCode(err, x.ErrorGenerateJWTToken, "failed generateToken for username %s", username)
	}

	return token, nil
}
