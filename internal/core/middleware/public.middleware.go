package middleware

import (
	"fmt"
	"net/http"
	"wild-workouts-go-ddd-example/internal/core/logger"
)

func PublicMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Add any common behavior for public routes here
        logger.Log(fmt.Sprintf("Public URL accessed: %s", r.URL.Path))
        next.ServeHTTP(w, r)
    })
}
