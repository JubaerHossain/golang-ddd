package userhttp

import (
	"net/http"

	"github.com/JubaerHossain/golang-ddd/internal/core/logger"
	"github.com/JubaerHossain/golang-ddd/internal/user/application"
	"github.com/JubaerHossain/golang-ddd/internal/user/domain/entity"
	"github.com/JubaerHossain/golang-ddd/pkg/utils"
	"go.uber.org/zap"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Implement GetUsers handler
	var users []*entity.ResponseUser
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
	// Implement CreateUser handler
	var newUser entity.User

	utils.BodyParse(&newUser, w, r, true) // Parse request body and validate it

	// Call the CreateUser function to create the user
	_, err := application.CreateUser(&newUser)
	if err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	// Write response
	utils.WriteJSONResponse(w, http.StatusCreated, map[string]interface{}{
		"message": "User created successfully",
	})
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Implement GetUserByID handler
	var user *entity.ResponseUser
	user, err := application.GetUser(r)
	if err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// Write response
	utils.WriteJSONResponse(w, http.StatusOK, map[string]interface{}{
		"message": "User fetched successfully",
		"users":   user,
	})

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Implement UpdateUser handler
	var newUser entity.UpdateUser
	utils.BodyParse(newUser, w, r, true) // Parse request body and validate it

	// Call the CreateUser function to create the user
	_, err := application.UpdateUser(r, &newUser)
	if err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Write response
	utils.WriteJSONResponse(w, http.StatusCreated, map[string]interface{}{
		"message": "User updated successfully",
	})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Implement DeleteUser handler
	err := application.DeleteUser(r)
	if err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// Write response
	utils.WriteJSONResponse(w, http.StatusOK, map[string]interface{}{
		"message": "User deleted successfully",
	})
}
