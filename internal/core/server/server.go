package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JubaerHossain/golang-ddd/internal/core/cache"
	"github.com/JubaerHossain/golang-ddd/internal/core/database"
	"github.com/JubaerHossain/golang-ddd/internal/core/logger"
	"github.com/JubaerHossain/golang-ddd/internal/core/routes"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

// Server represents the HTTP server
type Server struct {
	httpServer *http.Server
	cache      cache.CacheService // Include cache service
}

// NewServer creates a new instance of the Server
func NewServer() *Server {
	return &Server{}
}

// Start starts the HTTP server
func (s *Server) Start() error {
	// Initialize logger
	logger.Init()

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		logger.Error("Error loading .env file", zap.Error(err))
	}

	// Initialize database connection
	db, err := database.ConnectDB()
	if err != nil {
		logger.Error("Failed to connect to database", zap.Error(err))
		return err
	}
	if os.Getenv("MIGRATE") == "true" {
		database.MigrateDB(db)
	}

	// Initialize Redis cache service
	ctx := context.Background()
	cacheService, err := cache.NewRedisCacheService(ctx)
	if err != nil {
		logger.Error("Failed to initialize Redis cache service", zap.Error(err))
		return err
	}
	s.cache = cacheService

	// Get server address from environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not provided
	}
	addr := fmt.Sprintf(":%s", port)
	// make it beautiful with emoji
	fmt.Println("                   ")
	fmt.Println("                   ")
	fmt.Println("🚀 Server is starting  🚀   " + "http://localhost" + addr)
	fmt.Println("                  ")
	fmt.Println("                  ")
	// Register health check endpoint
	// Create HTTP server instance with middleware and routes
	router := routes.SetupRoutes()
	s.httpServer = &http.Server{
		Addr:    addr,
		Handler: router,
	}

	// Start the HTTP server in a separate goroutine
	go func() {
		logger.Info("Server is starting", zap.String("address", addr))
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Could not start server", zap.Error(err))
		}
	}()

	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	logger.Info("Server is shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		logger.Error("Could not gracefully shutdown server", zap.Error(err))
	}

	logger.Info("Server stopped")

	return nil
}
