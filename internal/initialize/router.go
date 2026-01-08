package initialize

import (
	"GolangBackendEcommerce/global"
	"GolangBackendEcommerce/internal/routers"

	"github.com/gin-gonic/gin"
)

// in golang, we must capitalize the function name to make it public
func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// middlewares
	r.Use() // logging
	r.Use() // cross
	r.Use() // global limitor

	manageRouter := routers.RouterGroupApp.Manager
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/v1/2025")
	{
		MainGroup.GET("/checkStatus", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}
	{
		manageRouter.InitUserRouter(MainGroup)
		manageRouter.InitAdminRouter(MainGroup)
	}
	return r
}
