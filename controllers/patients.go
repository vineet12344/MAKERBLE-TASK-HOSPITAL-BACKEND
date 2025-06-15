package controllers

import (
	"hospital-backend/initializers"
	"hospital-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllPatients(c *gin.Context) {
	var patients []models.Patient

	if err := initializers.DB.Find(&patients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch patients"})
		return
	}

	c.JSON(http.StatusOK, patients)
}

func CreatePatients(c *gin.Context) {
	var patient models.Patient

	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := initializers.DB.Create(&patient).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create patient"})
		return
	}

	c.JSON(http.StatusCreated, patient)
}

func GetPatientByID(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient

	if err := initializers.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Patient Not Found patient"})
		return
	}

	c.JSON(http.StatusOK, patient)
}

// UpdatePatient - receptionist only
func UpdatePatient(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient
	if err := initializers.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	initializers.DB.Save(&patient)
	c.JSON(http.StatusOK, patient)
}

// UpdateMedicalInfo - doctor only
func UpdateMedicalInfo(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient
	if err := initializers.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}
	var input struct {
		MedicalHistory string `json:"medical_history"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	patient.MedicalHistory = input.MedicalHistory
	initializers.DB.Save(&patient)
	c.JSON(http.StatusOK, patient)
}

// DeletePatient - receptionist only
func DeletePatient(c *gin.Context) {
	id := c.Param("id")
	if err := initializers.DB.Delete(&models.Patient{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete patient"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Patient deleted"})
}
