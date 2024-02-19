package utils

import (
	"encoding/json"
	"net/http"

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

func ReturnResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	response := Response{
		Success: statusCode >= 200 && statusCode < 300,
		Message: message,
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
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
		errors[err.Field()] = err.Field() + " is " + err.Tag() + " " + err.Param()
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
