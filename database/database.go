package database

import (
	"fmt"
	"go-todo-web/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB

func GetDatabase() *gorm.DB {
	if database == nil {

		// data source name
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=30s",
			config.GetConfig().MYSQL_USER,
			config.GetConfig().MYSQL_PASSWORD,
			config.GetConfig().MYSQL_HOST,
			config.GetConfig().MYSQL_DATABASE,
		)

		// connect to database
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		// if error panic
		if err != nil {
			panic("failed to connect database")
		}

		// set database
		database = db
	}

	return database
}
