package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-lab/controllers"
	"github.com/gin-lab/models"
)

func main() {
	models.InitialDatabase()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
		})
	})

	r.POST("/user", controllers.CreateUser)
	r.GET("/user/:id", controllers.GetUser)
	r.GET("/user", controllers.GetAllUser)
	r.PATCH("/user/:id", controllers.UpdateUser)
	r.DELETE("/user/:id", controllers.DeleteUser)
	r.PATCH("/user/:id/recover", controllers.RecoverUser)

	r.Run(":8000")
}
