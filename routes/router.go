package routes

import (
	"feed-me/controllers"
	"feed-me/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.LoadHTMLFiles("views/error.html")
	api := r.Group("/api")
	{
		landingpage := api.Group("/")
		{
			landingpage.GET("/", middleware.CheckJwt, controllers.HandleLandingPage(r))
			// landingpage.POST("/logout", func(c *gin.Context) {
			// 	c.SetCookie("auth", "", 3600*24, "", "", false, true)
			// 	c.JSON(200, "logged out")
			// 	c.Redirect(http.StatusFound, "/")
			// })
		}
		auth := api.Group("/auth")
		{
			auth.GET("/", controllers.HandleLoginPage(r))
			auth.POST("/", controllers.HandleLoginLogic(r))
			auth.GET("/verify", controllers.HandleVerify(r))
		}
		app := api.Group("/dashboard")
		{
			app.GET("/", controllers.HanldTestLayout(r))
		}
	}
}
