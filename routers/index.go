package routers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes add all routing list here automatically get main router
func RegisterRoutes(route *gin.Engine, db *gorm.DB, authMiddleware *jwt.GinJWTMiddleware) {
	// route.NoRoute(func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	// })
	route.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		// claims := jwt.ExtractClaims(c)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	// route.POST("/login", authMiddleware.LoginHandler)
	// route.GET("/refresh_token", authMiddleware.RefreshHandler)

	//Add All route
	AuthRoutes(route, db, authMiddleware)
	ArticleRoutes(route, db, authMiddleware)
	UserRoutes(route, db, authMiddleware)
}
