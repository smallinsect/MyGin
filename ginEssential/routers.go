package main

import (
	"github.com/gin-gonic/gin"
	"xuqiulin.com/mygin/ginEssential/controller"
	"xuqiulin.com/mygin/ginEssential/middleware"
)

// CollectRouter CollectRouter
func CollectRouter(r *gin.Engine) *gin.Engine {

	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	return r
}
