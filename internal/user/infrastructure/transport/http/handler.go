package userhttp

import (
	"encoding/json"
	"fmt"
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
	fmt.Println(err3)

	// errors := newUser.Validate()
	// if len(errors) > 0 {
	// 	// Write validation errors as response
	// 	utils.WriteJSONError(w, http.StatusBadRequest, errors)
	// 	return
	// }

	// if err := newUser.Validate(); err != nil {
	// 	errors := err.(validator.ValidationErrors)
	// 	utils.WriteJSONEValidation(w, http.StatusBadRequest, errors)
	// 	return
	// }

	// Call the CreateUser function to create the user
	user, err := application.CreateUser(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write response
	utils.WriteJSONResponse(w, http.StatusCreated, map[string]interface{}{
		"message": "User created successfully",
		"user":    user,
	})
}
