package initializers

import (
	"fmt"
	"hospital-backend/models"
	"log"
)

func SyncDB() {
	err := DB.AutoMigrate(&models.User{}, &models.Patient{})
	if err != nil {
		log.Fatal("error migrating DB", err)
	}

	fmt.Println("✅ DB Migration Complete")

}
