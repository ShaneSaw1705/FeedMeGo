package middleware

import (
	"feed-me/initializers"
	"feed-me/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CheckJwt(c *gin.Context) {
	tokenString, err := c.Cookie("auth")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Youre not authorized to access this content"})
		c.Abort()
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Youre not authorized to access this content"})
		c.Abort()
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(http.StatusUnauthorized, "Youre authorization has expired please log in again")
			c.Abort()
			return
		}
		var user models.UserModel
		initializers.DB.First(&user, claims["sub"])
		fmt.Println(claims["sub"])
		if user.ID == 0 {
			c.JSON(http.StatusBadRequest, "unable to find an asocciated user")
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, "Please sign in to view this content")
	}
}
