package repositories

import "project/models"

// UserRepository is an interface that defines the data operations.
type UserRepository interface {
    GetUserByID(id int) (*models.User, error)
}

// userRepository is a concrete implementation of the UserRepository interface.
type userRepository struct {
    // Here, you might have a DB connection, etc.
}

// NewUserRepository returns an instance of the concrete implementation of UserRepository.
func NewUserRepository() UserRepository {
    return &userRepository{}
}

// GetUserByID is a method to retrieve a user by their ID.
func (r *userRepository) GetUserByID(id int) (*models.User, error) {
    // Normally, you'd fetch data from a DB.
    user := &models.User{ID: id, Name: "John Doe", Age: 30}
    return user, nil
}
