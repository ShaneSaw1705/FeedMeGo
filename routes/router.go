package routes

import (
	"feed-me/controllers"
	"feed-me/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	landingpage := r.Group("/")
	{
		landingpage.GET("/", middleware.CheckJwt, controllers.HandleLandingPage(r))
		landingpage.POST("/", controllers.HandleTestToast(r))
	}
	auth := r.Group("/auth")
	{
		auth.GET("/", controllers.HandleLoginPage(r))
		auth.POST("/", controllers.HandleLoginLogic(r))
		auth.GET("/verify", controllers.HandleVerify(r))
	}
}
