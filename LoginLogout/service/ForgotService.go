package service

import (
	"jwt-authentication-golang/database"
	"jwt-authentication-golang/models"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdatePassword struct {
	Password string `json:"password"`
}

func Forgot(context *gin.Context) {
	log.Println("bug0")
	var request UpdatePassword
	var user models.User
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	log.Println("bug1")
	// check if email exists and password is correct
	record := database.Instance.Where("email = ?", context.Param("email")).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	log.Println("bug2")
	user.Password = request.Password

	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	log.Println("bug3")

	database.Instance.Model(&user).Updates(user)

	context.JSON(http.StatusOK, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
}
