package routes

import (
	"github.com/CallumBicknell/go-webServer/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	DefaultRoutes(router)
	AuthRoutes(router)

	return router
}

func DefaultRoutes(incomingRoutes *gin.Engine) {
	index := new(controllers.IndexController)
	incomingRoutes.GET("/", index.Home)
}

func AuthRoutes(incomingRoutes *gin.Engine) {
	index := new(controllers.IndexController)
	incomingRoutes.GET("/auth", index.Home)
	incomingRoutes.POST("/auth", index.Home)
	incomingRoutes.GET("/register", index.Home)
	incomingRoutes.POST("/register", index.Home)
}
