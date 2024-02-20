package repository

import "github.com/JubaerHossain/golang-ddd/internal/user/domain/entity"

// UserRepository defines methods for user data access
type UserRepository interface {
	GetAllUsers(queryValues map[string][]string) ([]*entity.ResponseUser, error)
	GetUserByID(userID uint) (*entity.User, error)
	GetUser(userID uint) (*entity.ResponseUser, error)
	CreateUser(user *entity.User) (*entity.User, error)
	UpdateUser(oldUser *entity.User, user *entity.UpdateUser) (*entity.User, error)
	DeleteUser(user *entity.User) error
	ChangePassword(oldUser *entity.User, user *entity.UserPasswordChange) error
	Login(loginUser *entity.LoginUser) (*entity.LoginUserResponse, error)
}


