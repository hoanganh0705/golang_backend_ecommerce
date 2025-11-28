package global

import (
	"GolangBackendEcommerce/pkg/logger"
	"GolangBackendEcommerce/pkg/settings"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
)
