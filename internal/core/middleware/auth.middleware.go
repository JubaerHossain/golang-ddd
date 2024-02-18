package middleware

import (
    "net/http"
    "wild-workouts-go-ddd-example/internal/core/logger"
)

func JWTAuthorizationMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Check JWT token in the request header
        // Implement your JWT verification logic here

        // If authorization fails, return unauthorized status
        // Otherwise, proceed to the next handler
        logger.Log("JWT authorization check")
        next.ServeHTTP(w, r)
    })
}
