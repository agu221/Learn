package router

import (
	"net/http"
	"sports_team_manager/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/teams", controllers.GetTeams)
	router.GET("/players", controllers.GetPlayers)
	router.POST("/players", controllers.PostPlayers)
	router.POST("/teams/register", controllers.RegisterPlayerToTeam)
	router.POST("/login", controllers.UserLoginHandler)
	router.POST("/register", controllers.UserRegisterHandler)
	return router
}
