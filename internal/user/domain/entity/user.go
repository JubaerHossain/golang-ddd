package entity

import (
	"bytes"
	"strings"
)

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
	ID        uint   `json:"id"`
	Username  string `json:"username" validate:"required,min=3,max=50"`
	Email     string `json:"email" validate:"required,email,max=100"`
	Password  string `json:"password" validate:"required,min=6,max=20"`
	Role      Role   `json:"role" gorm:"default:chef" validate:"required,oneof=admin manager waiter chef"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Status    Status `json:"status" gorm:"default:pending" validate:"required,oneof=active inactive deleted pending"`
}

func (u *User) Validate() string {
    var errMsg bytes.Buffer

    if u.Username == "" {
        errMsg.WriteString("Username is required\n")
    }
    if u.Email == "" {
        errMsg.WriteString("Email is required\n")
    }
    if u.Password == "" {
        errMsg.WriteString("Password is required\n")
    }
    if u.Role == "" {
        errMsg.WriteString("Role is required\n")
    }
    if u.Status == "" {
        errMsg.WriteString("Status is required\n")
    }

    return strings.TrimSpace(errMsg.String())
}
