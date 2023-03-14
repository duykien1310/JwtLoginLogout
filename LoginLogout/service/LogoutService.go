package service

import (
	"jwt-authentication-golang/database"

	"github.com/gin-gonic/gin"
)

func Logout(context *gin.Context) {

	redisClient := database.InitRedis()

	// Xóa access token khỏi Redis cache với key là "access_token"
	err := redisClient.Del("access_token").Err()
	if err != nil {
		panic(err)
	}
	// context.SetCookie("access_token", "", -1, "/", "localhost", false, true)
}
