package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func HandleLoginPage(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		templ, err := template.New("login").ParseFiles("views/login.html", "templates/base.html")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Error parsing html files"})
		}
		templ.ExecuteTemplate(c.Writer, "base", gin.H{"Title": "Login"})
	}
}

func HandleLoginLogic(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body struct {
			Email string `form:"email"`
		}
		err := c.Bind(&body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "failed to read body"})
			return
		}
		fmt.Println(body)
		r.LoadHTMLFiles("templates/toast.html")
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": body.Email,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})

		tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
		if err != nil {
			c.JSON(http.StatusFailedDependency, gin.H{
				"Error": "failed to create jwt token",
			})
		}
		fmt.Println(tokenString)
		c.HTML(200, "toast", gin.H{
			"message": "<uk-icon icon='rocket'></uk-icon> Magic link sent to: " + body.Email,
			"status":  "primary",
		})
	}
}
