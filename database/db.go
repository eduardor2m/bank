package database

import (
	"github.com/eduardor2m/bank/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("bank.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&models.Client{})
	DB.AutoMigrate(&models.Account{})
	DB.AutoMigrate(&models.Transaction{})

}
