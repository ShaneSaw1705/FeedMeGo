package controllers

import (
	"feed-me/initializers"
	"feed-me/models"
	"feed-me/services"
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
			"exp": time.Now().Add(time.Hour * 1).Unix(),
		})

		tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
		if err != nil {
			c.JSON(http.StatusFailedDependency, gin.H{
				"Error": "failed to create jwt token",
			})
			return
		}
		fmt.Println(tokenString)
		err = services.SendMagicLink(body.Email, tokenString)
		if err != nil {
			fmt.Println(err)
			c.HTML(200, "toast", gin.H{
				"message": "<uk-icon icon='rocket'></uk-icon> Magic failed to send to: " + body.Email,
				"status":  "primary",
			})
			return
		}
		c.HTML(200, "toast", gin.H{
			"message": "<uk-icon icon='rocket'></uk-icon> Magic link sent to: " + body.Email,
			"status":  "primary",
		})
	}
}

func HandleVerify(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Query("token")
		if tokenString == "" {
			c.JSON(http.StatusBadRequest, "Failed to load token")
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Ensure the signing method is what you expect
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// Return the secret key for validation
			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			// Redirect to the login page if token parsing fails
			c.JSON(http.StatusBadRequest, "Error parsing jwt tooken")
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// Check token expiration
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				c.JSON(http.StatusBadRequest, "This link is expired please try again")
				return
			}
			sub, ok := claims["sub"].(string)
			if !ok {
				c.JSON(http.StatusBadRequest, "Invalid token claims")
				return
			}
			var User models.UserModel
			initializers.DB.First(&User, "email = ?", sub)
			if User.ID == 0 {
				User.Email = sub
				initializers.DB.Create(&User)
			}
			//TODO: Set jwt cookie
		}
		c.JSON(200, "looks good to me :)")
		return
	}
}
