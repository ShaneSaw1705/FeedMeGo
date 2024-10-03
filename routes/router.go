package routes

import (
	"feed-me/controllers"
	"feed-me/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.GET("/", controllers.HandleLoginPage(r))
			auth.POST("/", controllers.HandleLoginLogic(r))
			auth.GET("/verify", controllers.HandleVerify(r))
		}
		feed := api.Group("/feed")
		{
			feed.GET("/:id", controllers.HandleFeedById(r))
			feed.POST("/create", middleware.CheckJwt, controllers.HandleCreateFeed(r))

		}
	}
}
