package routes

import (
	"hospital-backend/controllers"
	"hospital-backend/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Message": "HELLO SERVER RUNNING. Please use this link to test API: https://makerble-task-hospital-backend.onrender.com/swagger/index.html",
		})
	})
	api := router.Group("/api")

	api.POST("/signup", controllers.Signup)
	api.POST("/login", controllers.Login)

	patients := api.Group("/patients")
	patients.Use(middleware.RequireAuth())

	//protecetd routes // ? Common for both Receptionist and Doctor
	patients.GET("/", controllers.GetAllPatients)
	patients.GET("/:id", controllers.GetPatientByID)

	// ? receptionist only routes
	patients.POST("/", middleware.CheckRole("receptionist"), controllers.CreatePatients)
	patients.PUT("/:id", middleware.CheckRole("receptionist"), controllers.UpdatePatient)
	patients.DELETE("/:id", middleware.CheckRole("receptionist"), controllers.DeletePatient)

	// ? doctor only routes
	patients.PUT("/:id/medical", middleware.CheckRole("doctor"), controllers.UpdateMedicalInfo)

}
