package main

import (
	"fmt"
	"hospital-backend/initializers"
	"hospital-backend/routes"
	"log"
	"os"

	_ "hospital-backend/docs" // Swagger generated docs

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializers.LoadEnv()

	fmt.Println("✅ Env Variables loaded successfully")
	fmt.Printf("🔌 Connecting to DB → host: %s | dbname: %s\n", os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	initializers.ConnectDB()
	initializers.SyncDB()

	gin.SetMode(gin.ReleaseMode) // Use release mode in production
}

// @title Hospital Management Backend API
// @version 1.0
// @description Role-based backend for doctors and receptionists
// @host makerble-task-hospital-backend-production.up.railway.app
// @BasePath /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	fmt.Println("🚀 Starting Hospital Backend Server...")

	r := gin.Default()

	// Adding CORS
	r.Use(cors.Default())

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register app routes
	routes.RegisterRoutes(r)

	// Use Railway's PORT env var or fallback to 8080 locally
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("🌐 Listening on port:", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}
