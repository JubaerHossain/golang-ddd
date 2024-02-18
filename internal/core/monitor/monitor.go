package monitor

import (
    "net/http"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

// RegisterMetrics registers Prometheus metrics.
func RegisterMetrics() {
    // Define Prometheus metrics
    requestsTotal := prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "myapp_requests_total",
            Help: "Total number of HTTP requests processed",
        },
        []string{"method", "status"},
    )

    // Register metrics
    prometheus.MustRegister(requestsTotal)
}

// MetricsHandler returns an HTTP handler function that serves Prometheus metrics.
func MetricsHandler() http.Handler {
    return promhttp.Handler()
}
