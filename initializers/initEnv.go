package initializers

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading ENV Variables", err)
	}

	fmt.Println("✅ Env Varibles loaded Successfully")

}
