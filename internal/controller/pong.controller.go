package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

// use this to make our routers acknowledge which controller we are calling
func NewPongController() *PongController {
	return &PongController{}
}

func (pc *PongController) Pong(c *gin.Context) {
	fmt.Println("My middleware handler")
	name := c.Param("name")
	name1 := c.DefaultQuery("name1", "hoanganh")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"name1":   name1,
		"message": "pong" + name,
		"uid":     uid,
		"users":   []string{"hoanganh", "nguyenvana", "tranthi"},
	})
}
