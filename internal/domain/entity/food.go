// food.go
package entity

import (
	"github.com/go-playground/validator/v10"
)

// Food represents a food item.
type Food struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name" validate:"required"`
	Price    float64 `json:"price" validate:"required,gte=0"`
	Category string  `json:"category" validate:"required"`
}

// NewFood creates a new Food instance.
func NewFood(name string, price float64, category string) (*Food, error) {
	food := &Food{Name: name, Price: price, Category: category}

	// Validate the food entity
	if err := food.Validate(); err != nil {
		return nil, err
	}

	return food, nil
}

// Validate checks if the food entity is valid.
func (f *Food) Validate() error {
	validate := validator.New()
	if err := validate.Struct(f); err != nil {
		return err
	}
	return nil
}

// UpdateName updates the name of the food item.
func (f *Food) UpdateName(name string) error {
	f.Name = name
	return nil
}

// UpdatePrice updates the price of the food item.
func (f *Food) UpdatePrice(price float64) error {
	f.Price = price
	return nil
}

// UpdateCategory updates the category of the food item.
func (f *Food) UpdateCategory(category string) error {
	f.Category = category
	return nil
}
