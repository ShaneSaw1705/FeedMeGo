package routes

import (
	"feed-me/controllers"
	"feed-me/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	landingpage := r.Group("/")
	{
		landingpage.GET("/", middleware.CheckJwt, controllers.HandleLandingPage(r))
		landingpage.POST("/", controllers.HandleTestToast(r))
		landingpage.POST("/logout", func(c *gin.Context) {
			c.SetCookie("auth", "", 3600*24, "", "", false, true)
			c.Redirect(http.StatusFound, "/")
		})
	}
	auth := r.Group("/auth")
	{
		auth.GET("/", controllers.HandleLoginPage(r))
		auth.POST("/", controllers.HandleLoginLogic(r))
		auth.GET("/verify", controllers.HandleVerify(r))
	}
}
