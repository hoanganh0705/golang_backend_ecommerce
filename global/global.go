package global

import (
	"GolangBackendEcommerce/pkg/logger"
	"GolangBackendEcommerce/pkg/settings"
	"database/sql"

	"github.com/redis/go-redis/v9"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
	Mdbc   *sql.DB
	Rdb    *redis.Client
)
