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
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	router.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})
	router.GET("/user-dashboard", func(c *gin.Context) {
		role := "player"
		switch role {
		case "player":
			c.Redirect(http.StatusFound, "/dashboard/player")
		case "coach":
			c.Redirect(http.StatusFound, "/dashboard/coach")
		case "manager":
			c.Redirect(http.StatusFound, "/dashboard/manager")
		default:
			c.HTML(http.StatusOK, "dashboard_generic.html", nil)
		}
	})
	router.GET("/dashboard/player", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard_player.html", nil)
	})
	router.POST("/players", controllers.PostPlayers)
	router.POST("/users/login", controllers.UserLoginHandler)
	router.POST("/users/register", controllers.UserRegisterHandler)
	return router
}
