package dataSeed

import (
	"time"

	"github.com/JubaerHossain/golang-ddd/internal/category/domain/entity"
	"github.com/JubaerHossain/golang-ddd/internal/core/logger"
	utilQuery "github.com/JubaerHossain/golang-ddd/pkg/query"
	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
)

var roles = []entity.Role{entity.Admin, entity.Manager, entity.Waiter, entity.Chef}

// SeedUsers generates and inserts dummy user data into the database.
func SeedUsers(db *gorm.DB, numUsers int) error {
	if err := db.Exec("DELETE FROM users").Error; err != nil {
		return err
	}
	logger.Info("Deleted all users")
	password, _ := utilQuery.HashPassword("password")
	for _, role := range roles {
		for i := 0; i < numUsers; i++ {
			var user entity.User
			user.Username = faker.Username()
			user.Email = string(role) + "@example.com"
			user.Password = password
			user.Role = role
			user.CreatedAt = time.Now()
			user.UpdatedAt = time.Now()
			user.Status = entity.Active

			if err := db.Create(&user).Error; err != nil {
				return err
			}
		}
	}
	return nil
}
