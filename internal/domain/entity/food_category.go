// food_category.go
package entity

import (
	"github.com/go-playground/validator/v10"
)

// FoodCategory represents a category for food items.
type FoodCategory struct {
	ID   uint   `json:"id"`
	Name string `json:"name" validate:"required"`
}

// NewFoodCategory creates a new FoodCategory instance with the given name.
func NewFoodCategory(name string) (*FoodCategory, error) {
	category := &FoodCategory{Name: name}

	// Validate category entity
	if err := category.Validate(); err != nil {
		return nil, err
	}

	return category, nil
}

// Validate validates the food category entity.
func (fc *FoodCategory) Validate() error {
	validate := validator.New()
	return validate.Struct(fc)
}
