// category_service.go
package service

import (
	"github.com/JubaerHossain/golang-ddd/internal/domain/entity"
	"github.com/JubaerHossain/golang-ddd/internal/domain/repository"
	"github.com/go-playground/validator/v10"
)

// CategoryService represents the service for managing category-related operations.
type CategoryService struct {
	categoryRepo repository.CategoryRepository
	validate     *validator.Validate
}

// NewCategoryService creates a new instance of CategoryService.
func NewCategoryService(categoryRepo repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepo: categoryRepo,
		validate:     validator.New(),
	}
}

// GetCategoryByID retrieves a food category by its ID.
func (cs *CategoryService) GetCategoryByID(id uint) (*entity.FoodCategory, error) {
	category, err := cs.categoryRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

// GetAllCategories retrieves all food categories.
func (cs *CategoryService) GetAllCategories() ([]*entity.FoodCategory, error) {
	categories, err := cs.categoryRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// CreateCategory creates a new food category.
func (cs *CategoryService) CreateCategory(name string) (*entity.FoodCategory, error) {
	category := &entity.FoodCategory{Name: name}

	// Validate category entity
	if err := cs.validate.Struct(category); err != nil {
		return nil, err
	}

	// Persist category entity
	createdCategory, err := cs.categoryRepo.Create(category)
	if err != nil {
		return nil, err
	}

	return createdCategory, nil
}

// UpdateCategory updates an existing food category.
func (cs *CategoryService) UpdateCategory(id uint, name string) (*entity.FoodCategory, error) {
	category, err := cs.categoryRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update category properties
	category.Name = name

	// Validate updated category entity
	if err := cs.validate.Struct(category); err != nil {
		return nil, err
	}

	// Persist updated category entity
	updatedCategory, err := cs.categoryRepo.Update(category)
	if err != nil {
		return nil, err
	}

	return updatedCategory, nil
}

// DeleteCategory deletes a food category by its ID.
func (cs *CategoryService) DeleteCategory(id uint) error {
	err := cs.categoryRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
