package controllers

import (
	"feed-me/helpers"
	"feed-me/initializers"
	"feed-me/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleFeedById(c *gin.Context) {
	feedId := c.Param("id")
	if feedId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "bad url request"})
		return
	}
	user, err := helpers.GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Message": "There was an error fetching the user"})
		return
	}
	var feed models.Feed
	res := initializers.DB.First(&feed, "id = ?", feedId)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Message": "There was an error fetching the feed from the database"})
		return
	}
	if feed.AuthorId != int(user.ID) {
		c.JSON(http.StatusUnauthorized, gin.H{"Message": "You are not autherized to access this feed"})
		return
	}
	c.JSON(200, gin.H{"feed": feed})
}

func HandleUserFeeds(c *gin.Context) {
	user, err := helpers.GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "there was an error with the get address"})
		fmt.Printf("Error occured: %s", err)
		return
	}
	var feeds []models.Feed
	res := initializers.DB.Find(&feeds, "author_id = ?", user.ID)
	if res.Error != nil {
		c.JSON(http.StatusFailedDependency, gin.H{"Message": "failed to fetch from database"})
		return

	}
	fmt.Print(feeds)
	c.JSON(http.StatusOK, feeds)
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

func HandleDeleteFeed(c *gin.Context) {
	feedId := c.Param("id")
	if feedId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Error, param id not found"})
		return
	}

	user, err := helpers.GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Message": "You're not authorized to delete this feed"})
		return
	}

	var feed models.Feed
	res := initializers.DB.First(&feed, "id = ?", feedId)
	if res.Error != nil {
		if res.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"Message": "Feed not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": "Failed to fetch feed from database"})
		}
		return
	}

	if feed.AuthorId != int(user.ID) {
		c.JSON(http.StatusUnauthorized, gin.H{"Message": "You're not authorized to delete this feed"})
		return
	}

	res = initializers.DB.Delete(&feed)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "There was an error deleting the feed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Feed deleted successfully"})
}
