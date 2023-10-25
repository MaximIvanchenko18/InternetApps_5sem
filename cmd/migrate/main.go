package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"InternetApps_5sem/internal/app/ds"
	"InternetApps_5sem/internal/app/dsn"
)

func main() {
	_ = godotenv.Load()
	db, err := gorm.Open(postgres.Open(dsn.FromEnv()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(
		&ds.User{},
		&ds.Cargo{},
		&ds.Flight{},
		&ds.FlightCargo{},
	)
	if err != nil {
		panic("cant migrate db")
	}
}
