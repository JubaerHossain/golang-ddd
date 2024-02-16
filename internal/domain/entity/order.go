// order.go
package entity

import (
	"github.com/go-playground/validator/v10"
)

// OrderStatus represents the status of an order.
type OrderStatus string

const (
	// OrderStatusPending represents an order that is pending.
	OrderStatusPending OrderStatus = "pending"

	// OrderStatusProcessing represents an order that is being processed.
	OrderStatusProcessing OrderStatus = "processing"

	// OrderStatusCompleted represents an order that is completed.
	OrderStatusCompleted OrderStatus = "completed"

	// OrderStatusCancelled represents an order that is cancelled.
	OrderStatusCancelled OrderStatus = "cancelled"
)

// Order represents an order for food items.
type Order struct {
	ID           uint        `json:"id"`
	CustomerName string      `json:"customer_name" validate:"required"`
	Items        []*OrderItem `json:"items"`
	Status       OrderStatus `json:"status"`
}

// OrderItem represents an item in an order.
type OrderItem struct {
	FoodID   uint   `json:"food_id"`
	Quantity uint   `json:"quantity"`
}

// NewOrder creates a new Order instance with the given customer name.
func NewOrder(customerName string) *Order {
	return &Order{
		CustomerName: customerName,
		Status:       OrderStatusPending,
	}
}

// Validate validates the order entity.
func (o *Order) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}
