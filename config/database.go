package config

import (
	"fmt"

	"github.com/kurir-go/app/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database Connection is creating a new connection to our database
func DatabaseConnection() *gorm.DB {
	dbUser := "postgres"
	dbPass := "postgres"
	dbHost := "localhost"
	dbName := "kurir_go"
	dbPort := "5432"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	db.AutoMigrate(&model.User{}, &model.Shipment{}, &model.ShipmentHistory{})
	return db
}

func CloseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection database")
	}
	dbSQL.Close()
}
