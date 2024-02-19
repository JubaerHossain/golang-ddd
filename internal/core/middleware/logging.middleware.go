// File: wild-workouts-go-ddd-example/internal/core/middleware/logging.go

package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/JubaerHossain/golang-ddd/internal/core/logger"
	"go.uber.org/zap"
)

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        fmt.Println("LoggingMiddleware")

        // Call the next handler
        next.ServeHTTP(w, r)

        // Log request details after handling
        logger.Info("Request handled",
            zap.String("method", r.Method),
            zap.String("url", r.URL.Path),
            zap.Duration("duration", time.Since(start)),
        )
    })
}
