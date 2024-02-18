// File: core/database/database.go

package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/JubaerHossain/golang-ddd/internal/user/domain/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConnectDB initializes a connection to the PostgreSQL database
// ConnectDB initializes a connection to the PostgreSQL database
// ConnectDB initializes a connection to the PostgreSQL database
func ConnectDB() (*gorm.DB, error) {
	// Create database connection
	dsn := os.Getenv("DATABASE_URL") // Get database URL from environment variable
	if dsn == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable is not set")
	}
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // Disable prepared statements
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Configure connection pool settings
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database connection pool: %w", err)
	}
	sqlDB.SetMaxIdleConns(10)   // Maximum number of idle connections in the pool
	sqlDB.SetMaxOpenConns(100)  // Maximum number of open connections to the database
	sqlDB.SetConnMaxLifetime(0) // Maximum amount of time a connection may be reused

	// Test the database connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for i := 0; i < 3; i++ { // Retry logic
		err := sqlDB.PingContext(ctx)
		if err == nil {
			break
		}
		log.Printf("failed to ping database: %v (attempt %d)", err, i+1)
		time.Sleep(2 * time.Second) // Wait before retrying
	}
	if err != nil {
		return nil, fmt.Errorf("failed to ping database after multiple attempts: %w", err)
	}

	// defer func() {
	// 	dbInstance, _ := db.DB()
	// 	_ = dbInstance.Close()
	// }()

	log.Println("connected to database")

	return db, nil
}

// MigrateDB performs database migration
func MigrateDB(db *gorm.DB) error {
	// Add your database migration logic here
	// For example:
	if err := db.AutoMigrate(
		&entity.User{},
	); err != nil {
		return fmt.Errorf("failed to perform database migrations: %w", err)
	}
	log.Println("database migration completed")
	return nil
}

// CreateDB creates the database schema based on the provided models
func CreateDB(db *gorm.DB, models ...interface{}) error {
	err := db.AutoMigrate(models...)
	if err != nil {
		return fmt.Errorf("failed to create database schema: %w", err)
	}
	log.Println("database schema created")
	return nil
}
