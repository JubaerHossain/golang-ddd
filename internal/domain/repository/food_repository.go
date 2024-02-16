// food_repository.go
package repository

import "github.com/JubaerHossain/golang-ddd/internal/domain/entity"

// FoodRepository defines the interface for interacting with food data storage.
type FoodRepository interface {
	GetAll() ([]*entity.Food, error)
	GetByID(id uint) (*entity.Food, error)
	Create(food *entity.Food) (*entity.Food, error)
	Update(food *entity.Food) (*entity.Food, error)
	Delete(id uint) error
}


