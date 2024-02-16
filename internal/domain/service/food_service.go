// food_service.go
package service

import (
	"github.com/JubaerHossain/golang-ddd/internal/domain/entity"
	"github.com/JubaerHossain/golang-ddd/internal/domain/repository"
	"github.com/go-playground/validator/v10"
)

// FoodService represents the service for managing food-related operations.
type FoodService struct {
	foodRepo repository.FoodRepository
	validate *validator.Validate
}

// NewFoodService creates a new instance of FoodService.
func NewFoodService(foodRepo repository.FoodRepository) *FoodService {
	return &FoodService{
		foodRepo: foodRepo,
		validate: validator.New(),
	}
}

// GetAllFood retrieves all food items.
func (fs *FoodService) GetAllFood() ([]*entity.Food, error) {
	foods, err := fs.foodRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return foods, nil
}

// GetFoodByID retrieves a food item by its ID.
func (fs *FoodService) GetFoodByID(id uint) (*entity.Food, error) {
	food, err := fs.foodRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return food, nil
}

// CreateFood creates a new food item.
func (fs *FoodService) CreateFood(name string, price float64, category string) (*entity.Food, error) {
	food := &entity.Food{Name: name, Price: price, Category: category}

	// Validate food entity
	if err := fs.validate.Struct(food); err != nil {
		return nil, err
	}

	// Persist food entity
	createdFood, err := fs.foodRepo.Create(food)
	if err != nil {
		return nil, err
	}

	return createdFood, nil
}

// UpdateFood updates an existing food item.
func (fs *FoodService) UpdateFood(id uint, name string, price float64, category string) (*entity.Food, error) {
	food, err := fs.GetFoodByID(id)
	if err != nil {
		return nil, err
	}

	// Update food properties
	if name != "" {
		if err := food.UpdateName(name); err != nil {
			return nil, err
		}
	}
	if price > 0 {
		if err := food.UpdatePrice(price); err != nil {
			return nil, err
		}
	}
	if category != "" {
		if err := food.UpdateCategory(category); err != nil {
			return nil, err
		}
	}

	// Validate updated food entity
	if err := fs.validate.Struct(food); err != nil {
		return nil, err
	}

	// Persist updated food entity
	updatedFood, err := fs.foodRepo.Update(food)
	if err != nil {
		return nil, err
	}

	return updatedFood, nil
}

// DeleteFood deletes a food item by its ID.
func (fs *FoodService) DeleteFood(id uint) error {
	err := fs.foodRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
