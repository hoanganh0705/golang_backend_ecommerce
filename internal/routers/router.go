package router

import (
	c "GolangBackendEcommerce/internal/controller"

	"github.com/gin-gonic/gin"
)

// in golang, we must capitalize the function name to make it public

func NewRouter() *gin.Engine {
	r := gin.Default()

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
