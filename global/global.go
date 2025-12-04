package global

import (
	"GolangBackendEcommerce/pkg/logger"
	"GolangBackendEcommerce/pkg/settings"

	"gorm.io/gorm"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)
