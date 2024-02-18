package middleware

import (
	"net/http"
	"time"

	"github.com/JubaerHossain/golang-ddd/internal/core/logger"
	"go.uber.org/zap"
)

func JWTAuthorizationMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Check JWT token in the request header
        // Implement your JWT verification logic here

        // If authorization fails, return unauthorized status
        // Otherwise, proceed to the next handler
        
        next.ServeHTTP(w, r)
		start := time.Now()
		logger.Info("Request handled",
            zap.String("method", r.Method),
            zap.String("url", r.URL.Path),
            zap.Duration("duration", time.Since(start)),
        )
    })
}
