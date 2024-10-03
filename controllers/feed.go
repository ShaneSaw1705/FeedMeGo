package controllers

import (
	"feed-me/helpers"
	"feed-me/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleFeedById(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func HandleCreateFeed(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := helpers.GetCurrentUser(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": fmt.Sprintf("Error fetching current user: %s", err)})
			return
		}
		var responseBody struct {
			title string
		}

		err = c.Bind(&responseBody)
		if err != nil {
			c.JSON(http.StatusFailedDependency, gin.H{"Error": fmt.Sprintf("an error occured %s", err)})
			return
		}
		feed := models.Feed{
			Title:    responseBody.title,
			AuthorId: user.ID,
		}
	}
}
