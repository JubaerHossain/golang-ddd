// category_repository.go
package repository

import "github.com/JubaerHossain/golang-ddd/internal/domain/entity"

// CategoryRepository defines the interface for interacting with food category data storage.
type CategoryRepository interface {
	GetByID(id uint) (*entity.FoodCategory, error)
	GetAll() ([]*entity.FoodCategory, error)
	Create(category *entity.FoodCategory) (*entity.FoodCategory, error)
	Update(category *entity.FoodCategory) (*entity.FoodCategory, error)
	Delete(id uint) error
}