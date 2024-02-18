package entity

// User represents the user entity
type User struct {
    ID       uint   `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    // Add more fields as needed
}
