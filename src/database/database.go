package database

import (
	"adver_cart/src/environment"
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseInterface interface {
	NewConnection() *gorm.DB
}

type ConnectDB struct {
	DB *gorm.DB
}

//NewConnection
func NewConnection() *gorm.DB {

	env := environment.Load()
	dns := generateDNS(env.MySqlUsername,
		env.MySqlPassword,
		env.MySqlHost,
		env.MySqlPort,
		env.MySqlDatabaseName)

	mysql, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dns,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(fmt.Sprintf("❌ [mysql] failed to connect database: %s", err))
	}

	isDebug, err := strconv.ParseBool(os.Getenv("APP_DEBUG"))
	if err == nil && isDebug {
		mysql = mysql.Debug()
	}
	fmt.Println("✅ [mysql] connected " + env.MySqlDatabaseName)
	return mysql

}

// generateDNS for database
func generateDNS(user string, pass string, host string, port string, db string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		pass,
		host,
		port,
		db)
}
