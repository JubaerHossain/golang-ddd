package utils

import (
	"encoding/json"
	"net/http"
)

// Response represents a standardized JSON response format.
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
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

func WriteJSONEValidation(w http.ResponseWriter, statusCode int, errors interface{}) {
	response := Response{
		Success: false,
		Message: "Validation error",
		Data:    errors,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
