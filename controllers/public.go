package controllers

import (
	"hospital-backend/initializers"
	"hospital-backend/models"
	"hospital-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Signup godoc
// @Summary Register a new user (Doctor or Receptionist)
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body models.User true "User registration data"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /signup [post]
func Signup(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var user models.User
	var err error

	err = c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid Parameters",
			"log":     err.Error(),
			"Message": "All Fields are mandatory",
		})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Couldnot hash password",
		})
		return
	}

	user.Password = hashedPassword

	if err = initializers.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": " Could not insert into DB ",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login godoc
// @Summary Login for doctor or receptionist
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body object true "Login credentials (email & password only)"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /login [post]
func Login(c *gin.Context) {
	var input struct {
		Email    string
		Password string
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid credentials format!!!"})
		return
	}

	var user models.User

	if err := initializers.DB.Where("Email=?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Email or password"})
		return
	}

	if !utils.CheckPassword(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":    user.ID,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}
