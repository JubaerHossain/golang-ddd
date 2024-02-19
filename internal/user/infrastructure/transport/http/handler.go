package userhttp

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/JubaerHossain/golang-ddd/internal/core/cache"
	"github.com/JubaerHossain/golang-ddd/internal/core/logger"
	"github.com/JubaerHossain/golang-ddd/internal/user/application"
	"github.com/JubaerHossain/golang-ddd/internal/user/domain/entity"
	"github.com/JubaerHossain/golang-ddd/pkg/utils"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func GetUsers(w http.ResponseWriter, r *http.Request, cacheService cache.CacheService) {
	// Implement GetUsers handler
	var users []*entity.User
	users, err := application.GetUsers(r)
	if err != nil {
		logger.Error("Failed to fetch users", zap.Error(err))
		utils.WriteJSONError(w, http.StatusInternalServerError, "Failed to fetch users")
		return
	}

	// Write response
	utils.WriteJSONResponse(w, http.StatusOK, map[string]interface{}{
		"message": "Users fetched successfully",
		"users":   users,
	})
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
	validateErr := validate.Struct(newUser)
	if validateErr != nil {
		utils.WriteJSONEValidation(w, http.StatusBadRequest, validateErr.(validator.ValidationErrors))
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

func GetUserByID(w http.ResponseWriter, r *http.Request, cacheService cache.CacheService) {
	// Implement GetUserByID handler
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Implement UpdateUser handler
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Implement DeleteUser handler
}
