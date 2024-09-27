package controllers

import (
	"net/http"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

func HandleLandingPage(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		r.LoadHTMLFiles("templates/base.html", "views/index.html")
		templ, err := template.New("").ParseFiles("./views/index.html", "./templates/base.html")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Error parsing html files"})
		}
		templ.ExecuteTemplate(c.Writer, "base", gin.H{"Title": "Hello World"})
	}
}

func HandleTestToast(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		time.Sleep(2 * time.Second)
		r.LoadHTMLFiles("./templates/toast.html")
		c.HTML(200, "toast", gin.H{
			"message": "<uk-icon icon='rocket'></uk-icon>testing",
			"status":  "primary",
		})
	}
}
