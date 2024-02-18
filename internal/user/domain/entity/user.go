package entity

// User represents the user entity
type User struct {
    ID       uint   `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
    Role     string `json:"role"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
    Status   string `json:"status"`
}
