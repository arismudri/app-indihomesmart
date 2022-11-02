package routers

import (
	"app-indihomesmart/infra/database"
	"app-indihomesmart/infra/logger"
	"app-indihomesmart/routers/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SetupRoute() *gin.Engine {

	environment := viper.GetBool("DEBUG")
	if environment {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	authMiddleware, err := middleware.JwtMiddleware()

	if err != nil {
		logger.Fatalf("middleware JwtMiddleware() error: %s", err)
	}

	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		logger.Fatalf("authMiddleware.MiddlewareInit() error:" + errInit.Error())
	}

	allowedHosts := viper.GetString("ALLOWED_HOSTS")
	router := gin.New()
	router.SetTrustedProxies([]string{allowedHosts})
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())

	RegisterRoutes(router, database.DB, authMiddleware) //routes register

	return router
}
