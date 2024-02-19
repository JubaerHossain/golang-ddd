package entity


type LoginUser struct {
	Email    string `json:"email" validate:"required,email,max=100"`
	Password string `json:"password" validate:"required,min=6,max=20"`
}

type LoginUserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Status   Status `json:"status"`
	Token    string `json:"token"`
}