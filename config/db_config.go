package config

import (
	"github.com/Nezent/GoPark/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=anon dbname=park port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	db.AutoMigrate(&models.ParkingLot{},&models.Slots{},&models.Vehicle{})

	DB = db
}