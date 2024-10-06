package user

import model "go-service/model"

// GetUserByID is a method to retrieve a user by their ID.
func (r *userRepository) GetUserByID(id int) (*model.User, error) {
    return &model.User{ID: id, Name: "John Doe", Age: 30}, nil
}
