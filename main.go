package main

import (
	"feed-me/initializers"
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.Env()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		r.LoadHTMLFiles("templates/base.html", "views/index.html")
		templ, err := template.New("").ParseFiles("./views/index.html", "./templates/base.html")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err})
		}
		templ.ExecuteTemplate(c.Writer, "base", gin.H{"Title": "Hello World"})
	})

	r.POST("/", func(c *gin.Context) {
		time.Sleep(2 * time.Second)
		r.LoadHTMLFiles("./templates/toast.html")
		c.HTML(200, "toast", gin.H{
			"message": "<uk-icon icon='rocket'></uk-icon>testing",
			"status":  "primary",
		})
	})

	r.Run(os.Getenv("port"))
}
