package initialize

import (
	"GolangBackendEcommerce/global"
	"GolangBackendEcommerce/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
