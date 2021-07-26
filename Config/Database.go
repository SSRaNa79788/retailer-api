package Config

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB * gorm.DB

type DBConfig struct{
	Host,User,Password,DBName string
	Port int
}

func BuildDBConfig () *DBConfig{
	return &DBConfig{
		Host : "localhost",
		Port : 3306,
		User : "root",
		Password : "",
		DBName : "retailer_service_database",
	}
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
