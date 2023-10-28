package photocontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/timoothy21/task-5-pbi-btpns-TimothyTheophilusHartono/models"
)

func AddPhoto(c *gin.Context) {
	var photo models.Photo

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, _ := c.Get("user")

	photo.UserID = int(user.(models.User).ID)
	photo.User = user.(models.User)

	models.DB.Create(&photo)
	c.AbortWithStatusJSON(http.StatusAccepted, gin.H{
		"message": "photo added!",
		"photo":   photo,
	})
}

func ViewPhoto(c *gin.Context) {
	var photos []models.Photo
	result := models.DB.Preload("User").Find(&photos)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": result.Error})
		return
	}

	c.AbortWithStatusJSON(http.StatusAccepted, gin.H{
		"message": "photo added!",
		"photo":   photos,
	})
}
