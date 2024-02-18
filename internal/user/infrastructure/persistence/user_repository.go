package persistence

import (
	"github.com/JubaerHossain/golang-ddd/internal/core/database"
	"github.com/JubaerHossain/golang-ddd/internal/core/logger"
	"github.com/JubaerHossain/golang-ddd/internal/user/domain/entity"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	// Add fields for database connection or any other dependencies
	db *gorm.DB
}

// NewUserRepository returns a new instance of UserRepositoryImpl
func NewUserRepository() (*UserRepositoryImpl, error) {
	conn, err := database.ConnectDB()
	if err != nil {
		logger.Error("Failed to connect to database", zap.Error(err))
		return nil, err
	}
	return &UserRepositoryImpl{db: conn}, nil
}

// GetAllUsers returns all users from the database
func (r *UserRepositoryImpl) GetAllUsers() ([]*entity.User, error) {
	// Implement logic to get all users
	users := []*entity.User{}
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByID returns a user by ID from the database
func (r *UserRepositoryImpl) GetUserByID(userID uint) (*entity.User, error) {
	// Implement logic to get user by ID
	user := &entity.User{}
	if err := r.db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// CreateUser saves a new user to the database
func (r *UserRepositoryImpl) CreateUser(user *entity.User) (*entity.User, error) {
	// Implement logic to save user
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser updates a user in the database
func (r *UserRepositoryImpl) UpdateUser(user *entity.User) (*entity.User, error) {
	// Implement logic to update user
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// DeleteUser deletes a user from the database

func (r *UserRepositoryImpl) DeleteUser(userID uint) error {
	// Implement logic to delete user
	if err := r.db.Delete(&entity.User{}, userID).Error; err != nil {
		return err
	}
	return nil
}

func FilterUsers(queryValues map[string][]string) (users []*entity.User, err error) {
	// Implement logic to filter users
	return users, nil
}
