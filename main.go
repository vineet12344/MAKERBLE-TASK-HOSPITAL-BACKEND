package main

import (
	"fmt"
	"hospital-backend/initializers"
	"hospital-backend/routes"
	"log"
	"os"

	_ "hospital-backend/docs" // Important: for generated Swagger files

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
	initializers.SyncDB()
	gin.SetMode(gin.ReleaseMode)
}

// @title Hospital Management Backend API
// @version 1.0
// @description Role-based backend for doctors and receptionists
// @host localhost:8080
// @BasePath /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	fmt.Println("!!! Welcome to Hospital Backend !!!")

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
