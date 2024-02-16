// order_service.go
package service

import (
	"errors"

	"github.com/JubaerHossain/golang-ddd/internal/domain/entity"
	"github.com/JubaerHossain/golang-ddd/internal/domain/repository"
	"github.com/go-playground/validator/v10"
)

// OrderService represents the service for managing order-related operations.
type OrderService struct {
	orderRepo repository.OrderRepository
	validate  *validator.Validate
}

// NewOrderService creates a new instance of OrderService.
func NewOrderService(orderRepo repository.OrderRepository) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
		validate:  validator.New(),
	}
}

// GetOrderByID retrieves an order by its ID.
func (os *OrderService) GetOrderByID(id uint) (*entity.Order, error) {
	order, err := os.orderRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

// CreateOrder creates a new order.
func (os *OrderService) CreateOrder(customerName string) (*entity.Order, error) {
	order := &entity.Order{CustomerName: customerName}

	// Validate order entity
	if err := os.validate.Struct(order); err != nil {
		return nil, err
	}

	// Persist order entity
	createdOrder, err := os.orderRepo.Create(order)
	if err != nil {
		return nil, err
	}

	return createdOrder, nil
}

// AddFoodToOrder adds more food items to an existing order.
func (os *OrderService) AddFoodToOrder(orderID uint, foodID uint, quantity uint) (*entity.Order, error) {
	order, err := os.GetOrderByID(orderID)
	if err != nil {
		return nil, err
	}

	// Check if order status allows adding more food
	if order.Status != entity.OrderStatusPending {
		return nil, errors.New("cannot add food to order with status other than pending")
	}

	// Add food to order
	updatedOrder, err := os.orderRepo.AddMoreFood(order, foodID, quantity)
	if err != nil {
		return nil, err
	}

	return updatedOrder, nil
}

// UpdateOrderStatus updates the status of an order.
func (os *OrderService) UpdateOrderStatus(orderID uint, status string) (*entity.Order, error) {
	order, err := os.GetOrderByID(orderID)
	if err != nil {
		return nil, err
	}

	// Update order status
	updatedOrder, err := os.orderRepo.StatusUpdate(order, status)
	if err != nil {
		return nil, err
	}

	return updatedOrder, nil
}

// CancelOrder cancels an order by its ID.
func (os *OrderService) CancelOrder(id uint) error {
	err := os.orderRepo.Cancel(id)
	if err != nil {
		return err
	}
	return nil
}
