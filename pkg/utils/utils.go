package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Response represents a standardized JSON response format.
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

// WriteJSONResponse writes a JSON response with the specified status code.
func WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	response := Response{
		Success: statusCode >= 200 && statusCode < 300,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// WriteJSONError writes a JSON error response with the specified status code and message.
func WriteJSONError(w http.ResponseWriter, statusCode int, message string) {
	response := Response{
		Success: false,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func WriteJSONEValidation(w http.ResponseWriter, statusCode int, error interface{}) {
	errors := make(map[string]string)
	for _, err := range error.(validator.ValidationErrors) {
		if strings.ToLower(err.Field()) != "" {
			errors[strings.ToLower(err.Field())] = err.Field() + " is " + err.Tag()
		}
	}
	response := ErrorResponse{
		Success: false,
		Message: "Validation error",
		Errors:  errors,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func BodyParse(s interface{}, w http.ResponseWriter, r *http.Request, isValidation bool) error {
	fmt.Println("BodyParse")
	fmt.Println(s)
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(s)
	fmt.Println(err)
	if err != nil {
		WriteJSONError(w, http.StatusBadRequest, "Invalid JSON")
		return err
	}

	if isValidation {
		validate := validator.New()
		validateErr := validate.Struct(s)
		fmt.Println(validateErr)
		if validateErr != nil {
			fmt.Println("validation error")
			WriteJSONEValidation(w, http.StatusBadRequest, validateErr.(validator.ValidationErrors))
			return validateErr
		}
	}
	return nil
}
