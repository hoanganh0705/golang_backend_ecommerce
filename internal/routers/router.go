package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// in golang, we must capitalize the function name to make it public

func pong(c *gin.Context) {
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

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{

		v1.GET("/ping", pong)
		v1.PUT("/ping", pong)
		v1.PATCH("/ping", pong)
		v1.DELETE("/ping", pong)
		v1.HEAD("/ping", pong)
		v1.OPTIONS("/ping", pong)
	}

	v2 := r.Group("/v2")
	{

		v2.GET("/ping/:name", pong)
		v2.PUT("/ping", pong)
		v2.PATCH("/ping", pong)
		v2.DELETE("/ping", pong)
		v2.HEAD("/ping", pong)
		v2.OPTIONS("/ping", pong)
	}

	v3 := r.Group("/v3")
	{

		v3.GET("/ping/", pong)
		v3.PUT("/ping", pong)
		v3.PATCH("/ping", pong)
		v3.DELETE("/ping", pong)
		v3.HEAD("/ping", pong)
		v3.OPTIONS("/ping", pong)
	}

	return r
}
