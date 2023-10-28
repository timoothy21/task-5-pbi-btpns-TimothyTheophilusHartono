package usercontroller

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/timoothy21/task-5-pbi-btpns-TimothyTheophilusHartono/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed to hash password"})
		return
	}

	user.Password = string(hash)

	user.CreatedAt = time.Now()

	models.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"messages": "Failed to read body"})
	}

	var user models.User
	models.DB.First(&user, "email = ?", body.Email)

	if user.Id == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid email or password"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid email or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid to create token"})
		fmt.Println(err)
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)

	c.AbortWithStatusJSON(http.StatusAccepted, gin.H{"user": user, "token": tokenString})

}

func Update(c *gin.Context) {
	var user models.User
	userId := c.Param("userId")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user.UpdateAt = time.Now()

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

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"message": "i'm logged in",
		"user":    user,
	})
}
