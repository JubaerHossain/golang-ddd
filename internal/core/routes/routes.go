package routes

import (
	"net/http"

	"github.com/JubaerHossain/golang-ddd/internal/core/cache"
	"github.com/JubaerHossain/golang-ddd/internal/core/health_check"
	"github.com/JubaerHossain/golang-ddd/internal/core/middleware"
)

// SetupRoutes initializes and returns the HTTP router with all routes.
func SetupRoutes(cacheService cache.CacheService) *http.ServeMux {
	router := http.NewServeMux()

	// Register routes
	router.Handle("/health", middleware.LoggingMiddleware(http.HandlerFunc(health_check.HealthCheckHandler())))
	// Add more routes as needed

	return router
}
