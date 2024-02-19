package routes

import (
	"net/http"

	"github.com/JubaerHossain/golang-ddd/internal/core/cache"
	"github.com/JubaerHossain/golang-ddd/internal/core/health"
	"github.com/JubaerHossain/golang-ddd/internal/core/middleware"
	"github.com/JubaerHossain/golang-ddd/internal/core/monitor"
	userhttp "github.com/JubaerHossain/golang-ddd/internal/user/infrastructure/transport/http"
	"github.com/JubaerHossain/golang-ddd/pkg/utils"
)

// SetupRoutes initializes and returns the HTTP router with all routes.
func SetupRoutes(cacheService cache.CacheService) *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("/", middleware.LoggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.WriteJSONResponse(w, http.StatusOK, map[string]interface{}{ "message": "Welcome to the API" })
	})))

	// Register health check endpoint
	router.Handle("/health", middleware.LoggingMiddleware(http.HandlerFunc(health.HealthCheckHandler())))

	// Register monitoring endpoint
	router.Handle("/metrics", monitor.MetricsHandler())
	// Add more routes as needed
	router.Handle("/users", userhttp.SetupUserRoutes(cacheService))

	return router
}
