package seed

import (
	"fmt"

	"github.com/JubaerHossain/golang-ddd/internal/core/database"
	dataSeed "github.com/JubaerHossain/golang-ddd/internal/core/database/seed/data"
	"github.com/JubaerHossain/golang-ddd/internal/core/logger"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func NewSeed() error {
	logger.Init()
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		logger.Error("Error loading .env file", zap.Error(err))
	}
	// Connect to the database
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println("failed to connect to database:", err)
		return err
	}

	// Seed dummy users
	if err := dataSeed.SeedUsers(db, 1); err != nil {
		fmt.Println("failed to seed users:", err)
		return err
	}

	fmt.Println("Successfully seeded")
	return nil
}
