// order_repository.go
package repository

import "github.com/JubaerHossain/golang-ddd/internal/domain/entity"

// OrderRepository defines the interface for interacting with order data storage.
type OrderRepository interface {
	GetByID(id uint) (*entity.Order, error)
	Create(order *entity.Order) (*entity.Order, error)
	Update(order *entity.Order) (*entity.Order, error)
	AddMoreFood(order *entity.Order, foodID uint, quantity uint) (*entity.Order, error)
	StatusUpdate(order *entity.Order, status string) (*entity.Order, error)
	Cancel(id uint) error
}
