package middleware

import (
	"fmt"
	"net/http"

	"github.com/JubaerHossain/golang-ddd/internal/core/auth"
	"github.com/JubaerHossain/golang-ddd/internal/core/logger"
	"github.com/JubaerHossain/golang-ddd/pkg/utils"
	"go.uber.org/zap"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Check if the request is authenticated
		// If not, return an unauthorized response
		// Otherwise, call the next handler
		// For now, we will just call the next handler
		token := r.Header.Get("Authorization")
		if token == "" {
			utils.WriteJSONError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		_, err := auth.VerifyToken(token)
		if err != nil {
			logger.Error("Failed to verify token", zap.Error(err))
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized")
			return
		}
		next.ServeHTTP(w, r)

	})
}
