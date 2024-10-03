package routes

import (
	"feed-me/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.LoadHTMLFiles("views/error.html")
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.GET("/", controllers.HandleLoginPage(r))
			auth.POST("/", controllers.HandleLoginLogic(r))
			auth.GET("/verify", controllers.HandleVerify(r))
		}
	}
}
