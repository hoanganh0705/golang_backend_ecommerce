package initialize

import (
	"GolangBackendEcommerce/global"
	"GolangBackendEcommerce/internal/po"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error("mysql error", zap.Error(err))
		panic(err)
	}
}

func InitMySQL() {
	// Implementation for initializing MySQL goes here
	m := global.Config.Mysql
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.UserName, m.Password, m.Host, m.Dbname)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	checkErrorPanic(err, "MySQL connection failed")
	global.Logger.Info("MySQL connected successfully")
	global.Mdb = db

	// set pool
	// pool is used to control the number of connections to the database
	SetPool()

	// migrate tables
	migrateTables()
}

// InitMySQL().SetPool()
func SetPool() {
	m := global.Config.Mysql
	sqlDB, err := global.Mdb.DB()
	checkErrorPanic(err, "Set MySQL connection pool failed")
	if err != nil {
		fmt.Printf("mysql error: %s::", err)
	}
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
}

func migrateTables() {
	err := global.Mdb.AutoMigrate(&po.User{}, &po.Role{})

	if err != nil {
		fmt.Println("Migrating tables error:", err)
	}
}
