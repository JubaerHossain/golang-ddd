package userhttp

import (
	"net/http"
)

// SetupUserRoutes initializes and returns the HTTP router with user routes.
func SetupUserRoutes(router *http.ServeMux) *http.ServeMux {
	router.Handle("GET /users", http.HandlerFunc(GetUsers))
	router.Handle("POST /users", http.HandlerFunc(CreateUser))
	router.Handle("GET /users/show/{id}", http.HandlerFunc(GetUserByID))
	router.Handle("PUT /users/update/{id}", http.HandlerFunc(UpdateUser))
	router.Handle("DELETE /users/delete/{id}", http.HandlerFunc(DeleteUser))

	return router
}
