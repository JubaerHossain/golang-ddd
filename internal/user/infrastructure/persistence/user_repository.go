package persistence

import (
	"fmt"
	"time"

	"github.com/JubaerHossain/golang-ddd/internal/core/database"
	"github.com/JubaerHossain/golang-ddd/internal/core/logger"
	"github.com/JubaerHossain/golang-ddd/internal/user/domain/entity"
	utilQuery "github.com/JubaerHossain/golang-ddd/pkg/query"
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
func (r *UserRepositoryImpl) GetAllUsers(queryValues map[string][]string) ([]*entity.User, error) {
	// Implement logic to get all users
	users := []*entity.User{}
	query := r.FilterUsers(queryValues)                  // Filter
	paginate := utilQuery.Pagination(query, queryValues) // Pagination
	if err := paginate.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByID returns a user by ID from the database
func (r *UserRepositoryImpl) GetUserByID(userID uint) (*entity.User, error) {
	// Implement logic to get user by ID
	user := &entity.User{}
	if err := r.db.First(&user, userID).Error; err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

// CreateUser saves a new user to the database
func (r *UserRepositoryImpl) CreateUser(user *entity.User) (*entity.User, error) {
	// Implement logic to save user
	user.CreatedAt = string(time.Now().Format("2006-01-02 15:04:05"))
	user.UpdatedAt = string(time.Now().Format("2006-01-02 15:04:05"))
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser updates a user in the database
func (r *UserRepositoryImpl) UpdateUser(oldUser *entity.User, user *entity.User) (*entity.User, error) {
	// Implement logic to update user
	user.UpdatedAt = string(time.Now().Format("2006-01-02 15:04:05"))
	if err := r.db.Model(&oldUser).Updates(user).Error; err != nil {
		return nil, err
	}

	updateUser := &entity.User{}

	if err := r.db.First(&updateUser, oldUser.ID).Error; err != nil {
		return nil, err
	}
	return updateUser, nil
}

// DeleteUser deletes a user from the database

func (r *UserRepositoryImpl) DeleteUser(userID uint) error {
	// Implement logic to delete user
	if err := r.db.Delete(&entity.User{}, userID).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) FilterUsers(queryValues map[string][]string) *gorm.DB {
	// Construct base query
	query := r.db.Model(&entity.User{})

	// Filter by name
	if names, ok := queryValues["username"]; ok && len(names) > 0 {
		query = query.Where("username LIKE ?", "%"+names[0]+"%")
	}

	// Filter by email
	if emails, ok := queryValues["email"]; ok && len(emails) > 0 {
		query = query.Where("email LIKE ?", "%"+emails[0]+"%")
	}

	// Filter by status
	if statuses, ok := queryValues["status"]; ok && len(statuses) > 0 {
		query = query.Where("status IN (?)", statuses)
	}

	// Filter by role
	if roles, ok := queryValues["role"]; ok && len(roles) > 0 {
		query = query.Where("role IN (?)", roles)
	}

	// Filter by date
	if dates, ok := queryValues["date"]; ok && len(dates) > 0 {
		query = query.Where("created_at >= ?", dates[0])
	}

	// Filter by date range
	if dateRange, ok := queryValues["date_range"]; ok && len(dateRange) > 0 {
		query = query.Where("created_at BETWEEN ? AND ?", dateRange[0], dateRange[1])
	}
	// orderBy
	if conditions, ok := queryValues["orderBy"]; ok && len(conditions) > 0 {
		query = query.Order(conditions[0])

	} else {
		query = query.Order("created_at desc")
	}

	return query
}
