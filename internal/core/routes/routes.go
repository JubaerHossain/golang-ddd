package routes

import (
	"net/http"

	"github.com/JubaerHossain/golang-ddd/internal/core/cache"
	"github.com/JubaerHossain/golang-ddd/internal/core/health_check"
	"github.com/JubaerHossain/golang-ddd/internal/core/middleware"
	"github.com/JubaerHossain/golang-ddd/internal/core/monitoring"
)

// SetupRoutes initializes and returns the HTTP router with all routes.
func SetupRoutes(cacheService cache.CacheService) *http.ServeMux {
	router := http.NewServeMux()

	// Register health check endpoint
	router.Handle("/health", middleware.LoggingMiddleware(http.HandlerFunc(health_check.HealthCheckHandler())))

	// Register monitoring endpoint
	router.Handle("/metrics", monitoring.MetricsHandler())

	// Add more routes as needed

	return router
}
