package user

import (
	"go-service/model"

	"github.com/palantir/stacktrace"
)

// IsExistsUsernameAndPassword is a function to Check if there is user with username and password in database
func (r *userRepository) IsExistsUsernameAndPassword(username string, password string) (isExists bool, err error) {
	isExists = false
	var user model.User
	if err = r.db.Raw(SelectUserByUsernameAndPassword, username, password).Scan(&user).Error; err != nil {
		return false, stacktrace.Propagate(err, "Failed Run %s with parameters %s %s", SelectUserByUsernameAndPassword, username, password)
	}

	return user.ID != 0, nil
}
