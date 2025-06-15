package controllers

import (
	"hospital-backend/initializers"
	"hospital-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllPatients godoc
// @Summary Get all patient records
// @Tags Patients
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Patient
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /patients [get]
func GetAllPatients(c *gin.Context) {
	var patients []models.Patient

	if err := initializers.DB.Find(&patients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch patients"})
		return
	}

	c.JSON(http.StatusOK, patients)
}

// CreatePatients godoc
// @Summary Register a new patient
// @Tags Patients
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param patient body models.Patient true "Patient data"
// @Success 201 {object} models.Patient
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /patients [post]
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

// GetPatientByID godoc
// @Summary Get a patient by ID
// @Tags Patients
// @Produce json
// @Security BearerAuth
// @Param id path int true "Patient ID"
// @Success 200 {object} models.Patient
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /patients/{id} [get]
func GetPatientByID(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient

	if err := initializers.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Patient Not Found"})
		return
	}

	c.JSON(http.StatusOK, patient)
}

// UpdatePatient godoc
// @Summary Update an entire patient record
// @Tags Patients
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Patient ID"
// @Param patient body models.Patient true "Updated patient data"
// @Success 200 {object} models.Patient
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /patients/{id} [put]
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

// UpdateMedicalInfo godoc
// @Summary Update only medical history (Doctor only)
// @Tags Patients
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Patient ID"
// @Param data body object true "Medical history payload"
// @Success 200 {object} models.Patient
// @Failure 400 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /patients/{id}/medical [put]
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

// DeletePatient godoc
// @Summary Delete a patient record
// @Tags Patients
// @Security BearerAuth
// @Param id path int true "Patient ID"
// @Success 204 "No Content"
// @Failure 403 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /patients/{id} [delete]
func DeletePatient(c *gin.Context) {
	id := c.Param("id")
	if err := initializers.DB.Delete(&models.Patient{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete patient"})
		return
	}
	c.Status(http.StatusNoContent)
}
