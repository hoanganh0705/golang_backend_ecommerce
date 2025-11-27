package initialize

import (
	"GolangBackendEcommerce/global"
	"fmt"
)

func Run() {
	// load configuration
	LoadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading configuration mysql", m.UserName, m.Password, m.Dbname)
	// initialize logger
	InitLogger()

	// initialize MySQL
	InitMySQL()

	// initialize Redis
	InitRedis()

	// initialize router
	r := InitRouter()

	// start server
	r.Run()
}
