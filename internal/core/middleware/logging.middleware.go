package middleware

import (
    "fmt"
    "net/http"
    "time"
    "wild-workouts-go-ddd-example/internal/core/logger"
)

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        // Call the next handler
        next.ServeHTTP(w, r)

        // Log request details after handling
        logger.Log(fmt.Sprintf(
            "Method: %s, URL: %s, Duration: %s",
            r.Method,
            r.URL.Path,
            time.Since(start),
        ))
    })
}
