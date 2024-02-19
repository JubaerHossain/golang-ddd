package userhttp

import (
	"net/http"

	"github.com/JubaerHossain/golang-ddd/internal/core/cache")



// SetupUserRoutes initializes and returns the HTTP router with user routes.
func SetupUserRoutes(cacheService cache.CacheService) *http.ServeMux {
	router := http.NewServeMux()
	// Register user routes
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) { GetUsers(w, r, cacheService) })
	router.HandleFunc("GET /user/{id}", func(w http.ResponseWriter, r *http.Request) { GetUserByID(w, r, cacheService) })
	router.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) { CreateUser(w, r) })
	router.HandleFunc("PUT /user/{id}", func(w http.ResponseWriter, r *http.Request) { UpdateUser(w, r) })
	router.HandleFunc("DELETE /user/{id}", func(w http.ResponseWriter, r *http.Request) { DeleteUser(w, r) })

	return router
}