package routes

import (
	"feed-me/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	landingpage := r.Group("/")
	{
		landingpage.GET("/", controllers.HandleLandingPage(r))
		landingpage.POST("/", controllers.HandleTestToast(r))
	}
}
