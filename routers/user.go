package routers

import (
	"app-indihomesmart/controllers"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(route *gin.Engine, db *gorm.DB, authMiddleware *jwt.GinJWTMiddleware) {
	ctrl := controllers.Controller{DB: db}
	v1 := route.Group("user").Use(authMiddleware.MiddlewareFunc())
	{
		v1.GET("/", ctrl.UserGetList)
		v1.GET("/:id", ctrl.UserGetById)
		v1.POST("/", ctrl.UserAdd)
		v1.PUT("/:id", ctrl.UserEdit)
		v1.DELETE("/:id", ctrl.UserDelete)
	}
}
