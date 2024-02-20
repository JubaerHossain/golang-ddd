package entity

import "time"

// Status represents the status of a user
type Status string

const (
	// Active status
	Active Status = "active"
	// Inactive status
	Inactive Status = "inactive"
	// Deleted status
	Deleted Status = "deleted"

	// Pending status
	Pending Status = "pending"
)

// Role represents the role of a user
type Role string

const (
	// Admin role
	Admin   Role = "admin"
	Manager Role = "manager"
	Waiter  Role = "waiter"
	Chef    Role = "chef"
)

// User represents the user entity
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"` // Primary key
	Username  string    `json:"username" validate:"required,min=3,max=50"`
	Email     string    `json:"email" validate:"required,email,max=100"`
	Password  string    `json:"password" validate:"required,min=6,max=20"`
	Role      Role      `json:"role" gorm:"default:chef" validate:"required,oneof=admin manager waiter chef"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Status    Status    `json:"status" gorm:"default:pending" validate:"required,oneof=active inactive deleted pending"`
}



// updateUser represents the user update request
type UpdateUser struct {
	Username string `json:"username" validate:"omitempty,min=3,max=50"`
	Email    string `json:"email" validate:"omitempty,email,max=100"`
	Role     Role   `json:"role" gorm:"default:chef" validate:"omitempty,oneof=admin manager waiter chef"`
	Status   Status `json:"status" gorm:"default:pending" validate:"omitempty,oneof=active inactive deleted pending"`
}

type UserPasswordChange struct {
	OldPassword string `json:"old_password" validate:"required,min=6,max=20"`
	NewPassword string `json:"new_password" validate:"required,min=6,max=20"`
}

// responseUser represents the user response
type ResponseUser struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    Status    `json:"status"`
}

type AuthUser struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     Role   `json:"role"`
	Status   Status `json:"status"`
}
