package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to Database", err)
	}

	var dbName string

	err = DB.Raw("SELECT current_database()").Scan(&dbName).Error

	if err != nil {
		log.Fatal("Error Fetching DB Name", err)
	}

	fmt.Printf("✅ Connected to PostgreSQL via GORM !! DB name is: %s \n", dbName)

}
