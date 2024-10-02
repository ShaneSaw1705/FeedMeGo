package controllers

import (
	"feed-me/initializers"
	"feed-me/models"
	"feed-me/services"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func HandleLoginPage(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		r.LoadHTMLFiles("views/login.html", "templates/base.html")
		c.HTML(200, "base", gin.H{"Title": "Login"})
	}
}

func HandleLoginLogic(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body struct {
			Email string `json:"email"`
		}
		err := c.Bind(&body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "failed to read body"})
			return
		}
		fmt.Println(body)

		// Load the toast template
		r.LoadHTMLFiles("templates/toast.html")

		// Create a new JWT token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": body.Email,
			"exp": time.Now().Add(time.Hour * 1).Unix(),
		})

		// Sign the token
		tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
		if err != nil {
			c.JSON(http.StatusFailedDependency, gin.H{"Error": "failed to create jwt token"})
			return
		}
		fmt.Println(tokenString)

		// Send the email asynchronously
		go func() {
			err = services.SendMagicLink(body.Email, tokenString)
			if err != nil {
				// Log the error without attempting to send a response
				fmt.Printf("Failed to send magic link to %s: %v\n", body.Email, err)
			}
		}()

		// Respond immediately to the user
		c.JSON(200, gin.H{"message": "email has been sent to" + body.Email})
	}
}

func HandleVerify(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		r.LoadHTMLFiles("views/error.html")
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

			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, "Error parsing jwt tooken")
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// Check token expiration
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				c.HTML(http.StatusUnauthorized, "error", gin.H{"Error": "This link is expired please try again"})
				return
			}
			// bind jwt email
			sub, ok := claims["sub"].(string)
			if !ok {
				c.JSON(http.StatusBadRequest, "unable to assert email")
				return
			}
			// init User
			var User models.UserModel
			// check if it exists
			initializers.DB.Where("email = ?", sub).First(&User)
			// if not exists assign the email and create a database entry
			if User.ID == 0 {
				User.Email = sub
				res := initializers.DB.Create(&User)
				if res.Error != nil {
					c.JSON(http.StatusFailedDependency, "Failed to create database entry")
					return
				}
			}
			createAuthToken(User, c)
		}
		c.Redirect(http.StatusFound, os.Getenv("Frontend_Url")+"/")
		return
	}
}

func createAuthToken(user models.UserModel, c *gin.Context) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusFailedDependency, gin.H{
			"Error": "failed to create jwt token",
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("auth", tokenString, 3600*24, "", os.Getenv("Frontend_Url"), false, true)
}
