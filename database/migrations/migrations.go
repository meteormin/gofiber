package migrations

import (
	"gorm.io/gorm"
	"log"
)

// Migrate
// db entity 스키마에 맞춰 자동으로 migration
func Migrate(db *gorm.DB, dst ...interface{}) {
	log.Println("Auto Migrate...")
	err := db.AutoMigrate(dst...)

	if err != nil {
		log.Println("Failed Migration")
		log.Println(err)
	}

	log.Println("Success Migrate...")
}
