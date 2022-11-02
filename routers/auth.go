package routers

import (
	"app-indihomesmart/controllers"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoutes(route *gin.Engine, db *gorm.DB, authMiddleware *jwt.GinJWTMiddleware) {
	ctrl := controllers.Controller{DB: db}
	authRoute := route.Group("auth")
	authRoute.POST("/login", authMiddleware.LoginHandler)
	authRoute.POST("/register", ctrl.UserAdd)
	authRoute.GET("/refresh_token", authMiddleware.RefreshHandler)
	authRoute.Use(authMiddleware.MiddlewareFunc())
	{
		authRoute.GET("/logout", ctrl.Logout)
		authRoute.GET("/user-profile", ctrl.UserProfile)
	}
}
