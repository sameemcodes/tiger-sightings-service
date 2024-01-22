package durable

import (
	"fmt"
	"reflect"
	"time"

	"tigerhall-kittens/cmd/models"
	"tigerhall-kittens/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDb *gorm.DB

func InitMysqlDb() (*gorm.DB, error) {
	user := config.GetEnv("DB_USER", "root")
	password := config.GetEnv("DB_PASSWORD", "default_password")
	dbname := config.GetEnv("DB_NAME", "default_user_db")

	dbSource := fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, dbname)
	var err error
	var sqldb *gorm.DB

	for i := 0; i < 5; i++ {
		sqldb, err = gorm.Open(mysql.Open(dbSource), &gorm.Config{})
		if err == nil {
			break
		}
		fmt.Println("Failed to connect to database. Retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		return nil, err
	}

	sqlDB, err := sqldb.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(time.Minute * 5)

	ticker := time.NewTicker(time.Minute * 1)
	go func() {
		for range ticker.C {
			err := sqlDB.Ping()
			if err != nil {
				fmt.Println("Lost connection to database. Reconnecting...")
				sqldb, err = gorm.Open(mysql.Open(dbSource), &gorm.Config{})
				if err != nil {
					fmt.Println("Failed to reconnect to database.")
				} else {
					fmt.Println("Successfully reconnected to database.")
				}
			}
		}
	}()

	MysqlDb = sqldb
	return MysqlDb, nil
}

func CloseDbConnection() {
	db, err := MysqlDb.DB()
	if err != nil {
		fmt.Printf("Error retrieving DB instance for closing: %v\n", err)
		return
	}
	err = db.Close()
	if err != nil {
		fmt.Printf("Error while closing DB connection: %v\n", err)
	}
}

var modelsToMigrate = []interface{}{
	&models.User{},
	&models.Tiger{},
	&models.TigerSightingData{},
}

func AutoMigrateModels(db *gorm.DB, models []interface{}) {
	for _, model := range models {
		err := db.AutoMigrate(model)
		if err != nil {
			fmt.Printf("Error migrating %s model: %v\n", reflect.TypeOf(model).Elem().Name(), err)
		} else {
			fmt.Printf("%s model migration successful\n", reflect.TypeOf(model).Elem().Name())
		}
	}
}

func InitMysqlDbMigration() {
	AutoMigrateModels(MysqlDb, modelsToMigrate)
}
