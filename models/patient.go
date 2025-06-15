package models

import "gorm.io/gorm"

type Patient struct {
	gorm.Model
	Name           string `json:"name"`
	Email          string `json:"email" gorm:"unique"`
	Gender         string `json:"gender"`
	Contact        string `json:"contact"`
	MedicalHistory string `json:"medical_history"`
}
