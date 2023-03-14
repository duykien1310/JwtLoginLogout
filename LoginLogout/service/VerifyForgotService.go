package service

import (
	"jwt-authentication-golang/database"
	"jwt-authentication-golang/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VfCode struct {
	Password string `json:"password"`
	Code     string `json:"code"`
}

func VerifyForgot(context *gin.Context) {

	log.Println("bug0.")
	var request VfCode
	var user models.User
	var vcode models.VerifyCode
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	log.Println("bug1.")
	// check if email exists and password is correct
	record := database.Instance.Where("email = ?", context.Param("email")).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	rc := database.Instance.Last(&vcode)
	if rc.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": rc.Error.Error()})
		context.Abort()
		return
	}
	log.Println("bug2.")
	if request.Code != strconv.Itoa(vcode.Code) {
		context.JSON(http.StatusExpectationFailed, gin.H{"error": "VerifyCode not correct!"})
	}

	log.Println("bug3.")
	user.Password = request.Password

	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	log.Println("bug4.")

	database.Instance.Model(&user).Updates(user)
	// if err := database.Instance.Delete(&vcode).Error; err != nil {
	// 	context.AbortWithStatusJSON(500, gin.H{"message": "Failed to delete data"})
	// 	return
	// }

	// Xóa tất cả dữ liệu trong bảng
	if err := database.Instance.Exec("DELETE FROM verify_codes").Error; err != nil {
		context.JSON(500, gin.H{"error": "Failed to delete data"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
}
