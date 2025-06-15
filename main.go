package main

import (
	"fmt"
	"hospital-backend/initializers"
	"hospital-backend/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
	initializers.SyncDB()
	gin.SetMode(gin.ReleaseMode)
}




func main() {
	fmt.Println("!!! Welcome to Hospital Backend !!!")

	r := gin.Default()
	routes.RegisterRoutes(r)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	err := r.Run(":" + port)
	if err != nil {
		log.Fatal("Error Runnign the Server", err)
	}

	fmt.Println("Server running on port ", port)

}
