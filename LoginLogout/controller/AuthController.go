package controller

import (
	"jwt-authentication-golang/service"

	"github.com/gin-gonic/gin"
)

func AuthController() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/login", service.Login)
		api.POST("/user/register", service.RegisterUser)
		api.GET("/logout", service.Logout)
	}
	return router
}
