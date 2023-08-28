package models

import (
	"log"

	"github.com/mrmissx/stashbin-backend/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	uri := utils.GetEnv("DB_URI", "")
	if uri == "" {
		log.Fatal("DB_URI is not set")
	}

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{TranslateError: true})
	if err != nil {
		panic(err)
	}

	log.Println("Migrating database...")
	if err := db.AutoMigrate(&Document{}); err != nil {
		panic(err)
	}
	log.Println("Database migrated!")

	DB = db
}
