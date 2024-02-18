package server

import (
    "fmt"
    "net/http"
    "github.com/JubaerHossain/golang-ddd/internal/core/middleware"
)

type Server struct {
    httpServer      *http.Server
    publicMux       *http.ServeMux
    authorizedMux   *http.ServeMux
}

func NewServer() *Server {
    return &Server{
        publicMux:     http.NewServeMux(),
        authorizedMux: http.NewServeMux(),
    }
}

func (s *Server) Start() error {
    // Register public routes
    s.publicMux.HandleFunc("/", handler)

    // Apply middleware for public routes
    publicHandler := middleware.PublicMiddleware(s.publicMux)

    // Apply JWT authorization middleware for authorized routes
    authorizedHandler := middleware.JWTAuthorizationMiddleware(s.authorizedMux)

    // Create a multiplexer for all routes
    mux := http.NewServeMux()
    mux.Handle("/", publicHandler)
    mux.Handle("/api/", authorizedHandler) // Assuming API endpoints need authorization

    // Create HTTP server instance
    s.httpServer = &http.Server{
        Addr:    ":8080",
        Handler: middleware.LoggingMiddleware(mux),
    }

    // Start the HTTP server in a goroutine
    go func() {
        if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            panic(err)
        }
    }()

    fmt.Println("Server is running on port 8080...")
    return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        // Handle GET request (e.g., fetch user details)
    case http.MethodPost:
        // Handle POST request (e.g., create a new user)
    case http.MethodPut:
        // Handle PUT request (e.g., update user details)
    case http.MethodDelete:
        // Handle DELETE request (e.g., delete user)
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
        fmt.Fprintf(w, "Method not allowed")
    }
}
