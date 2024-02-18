// Package health_check provides functionality for health checks.
package health_check

import (
	"net/http"
)

// HealthCheckHandler returns an HTTP handler function for health checks.
func HealthCheckHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
}
