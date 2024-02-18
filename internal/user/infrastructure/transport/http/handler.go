package userhttp

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/JubaerHossain/golang-ddd/internal/core/logger"
	"github.com/JubaerHossain/golang-ddd/internal/user/application"
	"github.com/JubaerHossain/golang-ddd/internal/user/domain/entity"
	"github.com/JubaerHossain/golang-ddd/pkg/utils"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Implement GetUsers handler
	queryValues := r.URL.Query() //

	logger.Info("GetUsers handler called")

	logger.Info("GetUsers handler called", zap.String("queryValues", queryValues.Encode()))

	// users , err := application.GetUsers(queryValues)
	// if err != nil {
	// 	// Handle error
	// }

	// Write response
	utils.WriteJSONResponse(w, http.StatusOK, map[string]interface{}{
		"message": "Users fetched successfully",
		"users":   []interface{}{}, // Placeholder for user data
	})
}

type ApiError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "Field is required"
	// Add more custom error messages as needed
	default:
		return "Invalid value"
	}
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	var newUser entity.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		if err == io.EOF {
			utils.WriteJSONError(w, http.StatusBadRequest, "Empty request body")
		} else {
			utils.WriteJSONError(w, http.StatusBadRequest, "Invalid JSON")
		}
		return
	}

	// fmt.Println("newUser", newUser)
	validate := validator.New()
	err3 := validate.Struct(newUser)
	if err3 != nil {
		utils.WriteJSONEValidation(w, http.StatusBadRequest, err3.(validator.ValidationErrors))
		return
	}

	// Call the CreateUser function to create the user
	user, err := application.CreateUser(&newUser)
	if err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	// Write response
	utils.WriteJSONResponse(w, http.StatusCreated, map[string]interface{}{
		"message": "User created successfully",
		"user":    user,
	})
}
