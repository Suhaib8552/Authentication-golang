package routes

import (
	"github.com/gin-gonic/gin"

	"auth-go/middleware"
	"auth-go/route_handlers"
)

func RegisterRoutes(router *gin.Engine) {

	authApi := router.Group("/api/v1/auth")

	authApi.POST("/register", route_handlers.HandleRegister)

	authApi.POST("/login", route_handlers.HandleLogin)

	authApi.GET("/profile", middleware.JwtMiddleware(), route_handlers.Handleprofile)

	authApi.POST("/logout", middleware.JwtMiddleware(), route_handlers.HandleLogout)
}
