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

func HandleUserFeeds(c *gin.Context) {
	id := c.Param("id")
	user, err := helpers.GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "there was an error with the get address"})
		return
	}
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "there was an error with the get address"})
		return
	}
	var feeds []models.Feed
	res := initializers.DB.Find(&feeds, "AuthorId = ?", user.ID)
	if res.Error != nil {
		c.JSON(http.StatusFailedDependency, gin.H{"Message": "failed to fetch from database"})
		return

	}
	c.JSON(http.StatusOK, gin.H{"Feeds": feeds})
}

func HandleCreateFeed(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := helpers.GetCurrentUser(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": fmt.Sprintf("Error fetching current user: %s", err)})
			return
		}
		var ResponseBody struct {
			Title string `json:"title"`
		}

		err = c.Bind(&ResponseBody)
		if err != nil {
			c.JSON(http.StatusFailedDependency, gin.H{"Message": fmt.Sprintf("an error occured %s", err)})
			return
		}
		feed := models.Feed{
			Title:    ResponseBody.Title,
			AuthorId: int(user.ID),
		}
		res := initializers.DB.Create(&feed)
		if res.Error != nil {
			c.JSON(http.StatusFailedDependency, gin.H{"Message": fmt.Sprintf("There was an error creating the feed: %s", res.Error)})
			return
		}
		c.JSON(200, gin.H{"Message": "created feed sucessfully"})
	}
}
