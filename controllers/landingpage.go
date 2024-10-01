package controllers

import (
	"github.com/gin-gonic/gin"
)

func HandleLandingPage(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		r.LoadHTMLFiles("views/landingpage.html")
		c.HTML(200, "landingpage.html", nil)
	}
}
