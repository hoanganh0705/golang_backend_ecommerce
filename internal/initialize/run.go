package initialize

import (
	"GolangBackendEcommerce/global"
	"fmt"

	"go.uber.org/zap"
)

func Run() {
	// load configuration
	LoadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading configuration mysql", m.UserName, m.Password, m.Dbname)

	// initialize logger
	InitLogger()
	// test logger
	global.Logger.Info("Config log ok!", zap.String("oke", "success"))

	// initialize MySQL
	InitMySQL()

	// initialize Redis
	InitRedis()

	// initialize router
	r := InitRouter()

	// start server
	r.Run(fmt.Sprintf(":%d", global.Config.Server.Port))
}
