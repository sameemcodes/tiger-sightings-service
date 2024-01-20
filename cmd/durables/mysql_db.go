package durable

import (
	"fmt"

	"tigerhall-kittens/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDb *gorm.DB

func InitMysqlDb() (*gorm.DB, error) {
	// Get database credentials from environment variables
	user := config.GetEnv("DB_USER", "root")
	password := config.GetEnv("DB_PASSWORD", "default_password")
	dbname := config.GetEnv("DB_NAME", "default_user_db")

	// Create the database connection string
	dbSource := fmt.Sprintf("%s:%s@/%s", user, password, dbname)
	fmt.Println("dbSource", dbSource)

	// Open a connection to the database using GORM
	sqldb, err := gorm.Open(mysql.Open(dbSource), &gorm.Config{})
	fmt.Println("MysqlDb", sqldb)
	if err != nil {
		return nil, err
	}
	MysqlDb = sqldb
	return sqldb, nil
}

func CloseDbConnection() {
	db, err := MysqlDb.DB()
	if err != nil {
		panic(err)
	}
	db.Close()
}
