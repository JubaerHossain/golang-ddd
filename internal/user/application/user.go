// File: internal/user/application/user_service.go

package application

import (
	"context"

	"github.com/JubaerHossain/golang-ddd/internal/user/domain/entity"
	"github.com/JubaerHossain/golang-ddd/internal/user/infrastructure/persistence"
)

func GetUsers(ctx context.Context) ([]*entity.User, error) {
	// Call repository to get all users
	repo, err := persistence.NewUserRepository()
	if err != nil {
		return nil, err
	}

	users, err2 := repo.GetAllUsers()
	if err2 != nil {
		return nil, err2
	}
	return users, nil
}

// CreateUser creates a new user
func CreateUser(user *entity.User) (*entity.User, error) {
	// Add any validation or business logic here before creating the user

	repo, err := persistence.NewUserRepository()
	if err != nil {
		return nil, err
	}

	user, err2 := repo.CreateUser(user)
	if err2 != nil {
		return nil, err2
	}
	return user, nil
}

// GetUserByID retrieves a user by ID
func GetUserByID(ctx context.Context, id uint) (*entity.User, error) {
	// Call repository to get user by ID
	repo, err := persistence.NewUserRepository()
	if err != nil {
		return nil, err
	}

	user, err2 := repo.GetUserByID(id)
	if err2 != nil {
		return nil, err2
	}
	return user, nil
}

// UpdateUser updates an existing user
func UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	// Call repository to update user
	repo, err := persistence.NewUserRepository()
	if err != nil {
		return nil, err
	}

	user, err2 := repo.UpdateUser(user)
	if err2 != nil {
		return nil, err2
	}
	return user, nil
}

// DeleteUser deletes a user by ID
func DeleteUser(ctx context.Context, id uint) error {
	// Call repository to delete user
	repo, err := persistence.NewUserRepository()
	if err != nil {
		return err
	}

	err2 := repo.DeleteUser(id)
	if err2 != nil {
		return err2
	}

	return nil
}

func filterUser(ctx context.Context, status_id string) error {
	// Call repository to delete user
	_, err := persistence.NewUserRepository()
	if err != nil {
		return err
	}

	// err2 := repo.FilterUser(status_id)
	// if err2 != nil {
	// 	return err2
	// }

	return nil
}
