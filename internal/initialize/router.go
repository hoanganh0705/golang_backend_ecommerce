package initialize

import (
	c "GolangBackendEcommerce/internal/controller"
	"GolangBackendEcommerce/internal/middlewares"
	"fmt"

	"github.com/gin-gonic/gin"
)

// middleware trong golang phải return về handlerFunction
func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> AA")
		c.Next()
		fmt.Println("After --> AA")
	}

}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> BB")
		c.Next()
		fmt.Println("After --> BB")
	}

}

func CC(c *gin.Context) {
	fmt.Println("Before --> CC")
	c.Next()
	fmt.Println("After --> CC")

}

// in golang, we must capitalize the function name to make it public
func InitRouter() *gin.Engine {
	r := gin.Default()

	// use the middleware
	r.Use(middlewares.AuthenMiddleware(), AA(), BB())
	r.Use(CC)
	// the middleware will be executed in the order they are added, before AA -> before BB -> before CC -> after CC -> after BB -> after AA

	v1 := r.Group("/v1")
	{

		// do not use this way because it is not good, we should use struct to know which controller we are calling
		v1.GET("/ping", c.NewPongController().Pong)
		v1.GET("/user/1", c.NewUserController().GetUserById)
		// v1.PATCH("/ping", controller.Pong)
		// v1.DELETE("/ping", controller.Pong)
		// v1.HEAD("/ping", controller.Pong)
		// v1.OPTIONS("/ping", controller.Pong)
	}

	return r
}
