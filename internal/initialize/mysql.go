package initialize

import (
	"GolangBackendEcommerce/global"
	"GolangBackendEcommerce/internal/models"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
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
	genTableDAO()

	// migrate tables
	// migrateTables()
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

func genTableDAO() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/models",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// // gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(global.Mdb) // reuse your gorm db
	// g.GenerateAllTable()
	g.GenerateModel("go_crm_user")

	// // Generate basic type-safe DAO API for struct `model.User` following conventions
	// g.ApplyBasic(model.User{})

	// // Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	// g.ApplyInterface(func(Querier) {}, model.User{}, model.Company{})

	// Generate the code
	g.Execute()

}

func migrateTables() {
	err := global.Mdb.AutoMigrate(
		// &po.User{},
		// &po.Role{},
		&models.GoCrmUserV2{},
	)

	if err != nil {
		fmt.Println("Migrating tables error:", err)
	}
}
