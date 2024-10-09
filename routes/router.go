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
			feed.GET("/byid/:id", middleware.CheckJwt, controllers.HandleFeedById)
			feed.GET("/alluser", middleware.CheckJwt, controllers.HandleUserFeeds)
			feed.POST("/create", middleware.CheckJwt, controllers.HandleCreateFeed(r))
			feed.DELETE("/byid/:id", middleware.CheckJwt, controllers.HandleDeleteFeed)
		}
	}
}
