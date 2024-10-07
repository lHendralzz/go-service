package user

import (
	"go-service/model"

	"github.com/palantir/stacktrace"
)

// GetUserByEmail is a function to Check if there is user with email in database
func (r *userRepository) GetUserByEmail(email string) (user model.User, err error) {
	if err = r.db.Raw(SelectUserByEmail, email).Scan(&user).Error; err != nil {
		return user, stacktrace.Propagate(err, "Failed Run Query %s with parameters %s", SelectUserByEmail, email)
	}

	return user, nil
}
