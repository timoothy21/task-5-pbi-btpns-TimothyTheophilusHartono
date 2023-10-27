package main

import (
	"github.com/gin-gonic/gin"
	"github.com/timoothy21/task-5-pbi-btpns-TimothyTheophilusHartono/controllers/usercontroller"
	"github.com/timoothy21/task-5-pbi-btpns-TimothyTheophilusHartono/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.POST("/users/register", usercontroller.Register)
	r.POST("/users/login", usercontroller.Login)
	r.PUT("/users/:userId", usercontroller.Update)
	r.DELETE("/users/:userId", usercontroller.Delete)

	r.Run()
}
