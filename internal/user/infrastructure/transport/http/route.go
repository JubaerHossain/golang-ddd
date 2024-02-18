package userhttp

import (
	"net/http"

	"github.com/JubaerHossain/golang-ddd/internal/core/cache")



// SetupUserRoutes initializes and returns the HTTP router with user routes.
func SetupUserRoutes(cacheService cache.CacheService) *http.ServeMux {
	router := http.NewServeMux()
	// Register user routes
	router.HandleFunc("/users", GetUsers)
	// router.HandleFunc("/user/{id}", GetUserByID).Methods("GET")
	// router.HandleFunc("/user", SaveUser).Methods("POST")
	// router.HandleFunc("/user/{id}", UpdateUser).Methods("PUT")
	// router.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")

	// Add more user routes as needed

	return router
}
