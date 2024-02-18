package services

import (
    "github.com/JubaerHossain/golang-ddd/internal/user/domain/entity"
    "github.com/JubaerHossain/golang-ddd/internal/user/domain/repository"
)

// UserService handles user-related operations
type UserService struct {
    userRepository repository.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepository repository.UserRepository) *UserService {
    return &UserService{userRepository: userRepository}
}

// GetAllUsers retrieves all users
func (s *UserService) GetAllUsers() ([]*entity.User, error) {
    return s.userRepository.GetAllUsers()
}

// GetUserByID retrieves a user by ID
func (s *UserService) GetUserByID(userID uint) (*entity.User, error) {
    return s.userRepository.GetUserByID(userID)
}

// CreateUser creates a new user
func (s *UserService) CreateUser(user *entity.User) (*entity.User, error) {
    return s.userRepository.CreateUser(user)
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(user *entity.User) (*entity.User, error) {
    return s.userRepository.UpdateUser(user)
}

// DeleteUser deletes a user by ID
func (s *UserService) DeleteUser(userID uint) error {
    return s.userRepository.DeleteUser(userID)
}
