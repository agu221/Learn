package router

import (
	"sports_team_manager/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/teams", controllers.GetTeams)
	router.GET("/players", controllers.GetPlayers)
	router.POST("/players", controllers.PostPlayers)

	return router
}
