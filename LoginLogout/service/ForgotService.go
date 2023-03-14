package service

import (
	"jwt-authentication-golang/database"
	"jwt-authentication-golang/middlewares"
	"jwt-authentication-golang/models"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestEmail struct {
	Email string `json:"email"`
}

func Forgot(context *gin.Context) {
	log.Println("bug0")
	var request RequestEmail
	var user models.User
	var verifyCode models.VerifyCode
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	log.Println("bug1")
	// check if email exists and password is correct

	record := database.Instance.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	random := middlewares.GoMail(request.Email)
	log.Println(request.Email)
	if random == -1 {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "don't send email"})
		context.Abort()
		return
	}
	verifyCode.Code = random
	database.Instance.Create(&verifyCode)
	context.JSON(http.StatusOK, gin.H{"status": "send email success!"})

	// log.Println("bug2")
	// user.Password = request.Password

	// if err := user.HashPassword(user.Password); err != nil {
	// 	context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	context.Abort()
	// 	return
	// }

	// log.Println("bug3")

	// database.Instance.Model(&user).Updates(user)

	// context.JSON(http.StatusOK, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
}
