package user

import (
	"fmt"
	"go-service/stdlib/auth"
	x "go-service/stdlib/error"
	"go-service/stdlib/hash"

	"github.com/palantir/stacktrace"
)

func (s *userService) Login(email string, password string) (string, error) {
	user, err := s.userRepository.GetUserByEmail(email)
	if err != nil {
		return "", stacktrace.PropagateWithCode(err, x.ErrorQuery, "Select User With Username = %s", email)
	}

	isValid := hash.ComparePassword(user.Password, password)
	if !isValid {
		return "", stacktrace.NewErrorWithCode(x.ErrorLogin, "email or password is invalid")
	}

	token, err := auth.GenerateToken(email, user.ID, s.opt.JWTKey)
	if err != nil {
		fmt.Println(err)
		return "", stacktrace.PropagateWithCode(err, x.ErrorGenerateJWTToken, "failed generateToken for email %s", email)
	}

	return token, nil
}
