package controllers

import (
	"feed-me/helpers"
	"feed-me/initializers"
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
		var ResponseBody struct {
			Title string `json:"title"`
		}

		err = c.Bind(&ResponseBody)
		if err != nil {
			c.JSON(http.StatusFailedDependency, gin.H{"Error": fmt.Sprintf("an error occured %s", err)})
			return
		}
		feed := models.Feed{
			Title:    ResponseBody.Title,
			AuthorId: int(user.ID),
		}
		res := initializers.DB.Create(&feed)
		if res.Error != nil {
			c.JSON(http.StatusFailedDependency, gin.H{"Error": fmt.Sprintf("There was an error creating the feed: %s", res.Error)})
			return
		}
		c.JSON(200, gin.H{"Message": "created feed sucessfully"})
	}
}
