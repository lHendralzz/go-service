package user

import model "go-service/model"

// GetUserProfile fetches a user profile based on business logic.
func (s *userService) GetUserProfile(id int) (*model.User, error) {
    return s.UserRepository.GetUserByID(id)
}