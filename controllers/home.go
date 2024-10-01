package controllers

import (
	"github.com/gin-gonic/gin"
)

func HanldTestLayout(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		r.LoadHTMLFiles("views/index.html", "templates/base.html")
		c.HTML(200, "base", gin.H{
			"Title": "Dashboard",
		})
	}
}
