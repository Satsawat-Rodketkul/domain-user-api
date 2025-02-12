package repository

import (
	"log"

	"github.com/Satsawat-Rodketkul/domain-user-api/internal/database/models"
	"gorm.io/gorm"
)

func register(db *gorm.DB, model models.RegisterEntity) {

	result := db.Create(model)

	if result.Error != nil {
		log.Fatal("Failed to create user profile:", result.Error)
	}

	log.Print("Create user profile success")
}
