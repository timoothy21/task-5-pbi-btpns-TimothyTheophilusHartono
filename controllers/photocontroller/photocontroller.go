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

func UpdatePhoto(c *gin.Context) {
	var photo models.Photo
	photoId := c.Param("photoId")

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, _ := c.Get("user")

	if models.DB.Model(&photo).Where("user_id = ?", user.(models.User).ID).Where("id = ?", photoId).Updates(&photo).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "data can't be updated"})
		return
	}

	c.AbortWithStatusJSON(http.StatusAccepted, gin.H{
		"message": "photo updated!",
	})
}

func DeletePhoto(c *gin.Context) {
	var photo models.Photo
	photoId := c.Param("photoId")

	user, _ := c.Get("user")

	if models.DB.Model(&photo).Where("user_id = ?", user.(models.User).ID).Delete(&photo, photoId).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data can't deleted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data has been deleted"})
}
