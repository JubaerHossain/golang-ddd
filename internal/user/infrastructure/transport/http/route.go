package userhttp

import (
	"net/http"

	"github.com/JubaerHossain/golang-ddd/internal/core/middleware"
)

// SetupUserRoutes initializes and returns the HTTP router with user routes.
func SetupUserRoutes(router *http.ServeMux) *http.ServeMux {


	// auth routes
	router.Handle("POST /auth/login", middleware.LimiterMiddleware(middleware.LoggingMiddleware(http.HandlerFunc(Login))))

	router.Handle("GET /users", middleware.LimiterMiddleware(middleware.Authenticate(http.HandlerFunc(GetUsers))))
	router.Handle("POST /users", middleware.LimiterMiddleware(middleware.Authenticate(http.HandlerFunc(CreateUser))))
	router.Handle("GET /users/{id}", middleware.LimiterMiddleware(middleware.Authenticate(http.HandlerFunc(GetUserByID))))
	router.Handle("PUT /users/{id}", middleware.LimiterMiddleware(middleware.Authenticate(http.HandlerFunc(UpdateUser))))
	router.Handle("DELETE /users/{id}", middleware.LimiterMiddleware(middleware.Authenticate(http.HandlerFunc(DeleteUser))))
	router.Handle("PUT /users/password-change/{id}", middleware.LimiterMiddleware(middleware.Authenticate(http.HandlerFunc(ChangePassword))))



	return router
}
