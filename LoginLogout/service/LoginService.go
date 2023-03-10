package service

import (
	"jwt-authentication-golang/database"
	"jwt-authentication-golang/models"
	"jwt-authentication-golang/util"

	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(context *gin.Context) {
	var request LoginRequest
	var user models.User
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	// check if email exists and password is correct
	record := database.Instance.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}
	tokenString, err := util.GenerateJWT(user.Email, user.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	redisClient := database.InitRedis()

	// Thêm access token vào Redis cache với key là "access_token" và giá trị là "my_access_token"
	errredis := redisClient.Set("access_token", tokenString, 0).Err()
	if errredis != nil {
		panic(errredis)
	}

	// context.SetCookie("access_token", tokenString, 10000, "/", "localhost", false, true)

	context.JSON(http.StatusOK, gin.H{"token": tokenString, "status": "success", "email": request.Email})
}
