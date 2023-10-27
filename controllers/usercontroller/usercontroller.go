package usercontroller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/timoothy21/task-5-pbi-btpns-TimothyTheophilusHartono/models"
)

func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	user.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	models.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func Login(c *gin.Context) {

}

func Update(c *gin.Context) {
	var user models.User
	userId := c.Param("userId")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user.UpdateAt = time.Now().Format("2006-01-02 15:04:05")

	if models.DB.Model(&user).Where("id = ?", userId).Updates(&user).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "data not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil diperbarui"})
}

func Delete(c *gin.Context) {
	var user models.User
	userId := c.Param("userId")

	if models.DB.Model(&user).Delete(&user, userId).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data can't deleted"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "data has been deleted"})
}
