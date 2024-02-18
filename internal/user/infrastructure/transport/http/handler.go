package userhttp

import (
	"net/http"

	"github.com/JubaerHossain/golang-ddd/internal/core/logger"
	"github.com/JubaerHossain/golang-ddd/pkg/utils"
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
