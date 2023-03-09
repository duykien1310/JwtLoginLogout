package service

import "github.com/gin-gonic/gin"

func Logout(context *gin.Context) {
	context.SetCookie("access_token", "", -1, "/", "localhost", false, true)
}
