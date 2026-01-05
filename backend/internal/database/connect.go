package database

import (
	"log"

	"github.com/Mywallet/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("Wallet.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	log.Println("Connect database successfully")

	err = DB.AutoMigrate(
		&models.Wallet{},
		&models.Category{},
		&models.Transaction{},
		&models.Budget{},
		&models.SavingsGoal{},
	)
	if err != nil {
		log.Fatal("migrate failed", err)
	}
}
