package migrations

import (
	"github.com/miniyus/gofiber/entity"
	"gorm.io/gorm"
	"log"
)

// Migrate
// db entity 스키마에 맞춰 자동으로 migration
func Migrate(db *gorm.DB) {
	log.Println("Auto Migrate...")
	err := db.AutoMigrate(
		&entity.AccessToken{},
		&entity.Action{},
		&entity.Permission{},
		&entity.Group{},
		&entity.User{},
		&entity.JobHistory{},
	)

	if err != nil {
		log.Fatalf("Failed Auto Migration")
	}

	log.Println("Success Auto Migrate...")
}
