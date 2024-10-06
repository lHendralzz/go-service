package user

import (
	"fmt"
)

// IsExistsUsernameAndPassword is a function to Check if there is user with username and password in database
func (r *userRepository)IsExistsUsernameAndPassword(username string, password string)(isExists bool, err error){
    isExists = false
    fmt.Println(SelectUserByUsernameAndPassword, username, password)
    return
}