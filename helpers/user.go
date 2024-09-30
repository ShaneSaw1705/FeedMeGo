package helpers

import (
	"errors"
	"feed-me/models"

	"github.com/gin-gonic/gin"
)

func GetCurrentUser(c *gin.Context) (models.UserModel, error) {
	value, exists := c.Get("user")
	if !exists {
		return models.UserModel{}, errors.New("the requested cookie does not exist")
	}

	user, ok := value.(models.UserModel)
	if !ok {
		return models.UserModel{}, errors.New("failed to conform the user cookie to the user model")
	}
	return user, nil
}
