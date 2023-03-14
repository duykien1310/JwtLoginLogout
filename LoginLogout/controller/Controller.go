package controller

import (
	"jwt-authentication-golang/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	// Initialize Router
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/login", service.Login)
		api.POST("/user/register", service.RegisterUser)
		api.GET("/logout", service.Logout)
		api.POST("/forgotpass", service.Forgot)
		api.PATCH("/changepass/:email", service.VerifyForgot)
	}
	return router
}
