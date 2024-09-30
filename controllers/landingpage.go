package controllers

import (
	"feed-me/models"
	"net/http"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

func HandleLandingPage(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		value, _ := c.Get("user")
		user, _ := value.(models.UserModel)

		templ, err := template.New("landingpage").ParseFiles("views/index.html", "templates/base.html")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Error parsing html files"})
		}
		templ.ExecuteTemplate(c.Writer, "base", gin.H{
			"Title":     "Home",
			"useremail": user.Email,
		})
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
