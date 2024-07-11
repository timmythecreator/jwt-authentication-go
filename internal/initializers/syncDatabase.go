package initializers

import (
	"go-jwt/internal/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
