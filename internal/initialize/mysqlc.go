package initialize

import (
	"GolangBackendEcommerce/global"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

func checkErrorPanicC(err error, errString string) {
	if err != nil {
		global.Logger.Error("mysql error", zap.Error(err))
		panic(err)
	}
}

func InitMySQLC() {
	// Implementation for initializing MySQL goes here
	m := global.Config.Mysql
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.UserName, m.Password, m.Host, m.Dbname)
	db, err := sql.Open("mysql", s)
	checkErrorPanicC(err, "MySQL connection failed")
	global.Logger.Info("MySQL connected successfully")
	global.Mdbc = db

	// set pool
	// pool is used to control the number of connections to the database
	SetPoolC()
	// genTableDAOC() // Disable code generation at runtime - should be done during development only

	// migrate tables
	// migrateTablesC()
}

// SetPoolC sets connection pool settings for MySQL
func SetPoolC() {
	m := global.Config.Mysql
	global.Mdbc.SetMaxIdleConns(m.MaxIdleConns)
	global.Mdbc.SetMaxOpenConns(m.MaxOpenConns)
	global.Mdbc.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
}
