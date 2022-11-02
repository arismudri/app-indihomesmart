package routers

import (
	"app-indihomesmart/controllers"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ArticleRoutes(route *gin.Engine, db *gorm.DB, authMiddleware *jwt.GinJWTMiddleware) {
	ctrl := controllers.Controller{DB: db}
	v1 := route.Group("article").Use(authMiddleware.MiddlewareFunc())
	{
		v1.GET("/", ctrl.ArticleGetList)
		v1.GET("/:id", ctrl.ArticleGetById)
		v1.POST("/", ctrl.ArticleAdd)
		v1.PUT("/:id", ctrl.ArticleEdit)
		v1.DELETE("/:id", ctrl.ArticleDelete)
	}
}
