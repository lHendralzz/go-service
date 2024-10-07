package user

import (
	"go-service/model"

	"github.com/palantir/stacktrace"
)

// IsExistsUsernameAndPassword is a function to Check if there is user with username and password in database
func (r *userRepository) GetPasswordByUsername(username string) (user model.User, err error) {
	if err = r.db.Raw(SelectUserByUsername, username).Scan(&user).Error; err != nil {
		return user, stacktrace.Propagate(err, "Failed Run Query %s with parameters %s", SelectUserByUsername, username)
	}

	return user, nil
}
