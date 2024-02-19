package repository

import "github.com/JubaerHossain/golang-ddd/internal/user/domain/entity"

// UserRepository defines methods for user data access
type UserRepository interface {
    GetAllUsers() ([]*entity.User, error)
    GetUserByID(userID uint) (*entity.User, error)
    CreateUser(user *entity.User) (*entity.User, error)
    UpdateUser(user *entity.User) (*entity.User, error)
    DeleteUser(userID uint) error
}


