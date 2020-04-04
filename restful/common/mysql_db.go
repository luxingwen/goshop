package common

import (
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"goshop/restful/config"
)

type Database struct {
	*gorm.DB
}

var db *gorm.DB

// Opening a database and save the reference to `Database` struct.
func Init() *gorm.DB {
	log.Info("init db")
	server := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", config.MySqlConf.UserName, config.MySqlConf.PassWd, config.MySqlConf.Host, config.MySqlConf.Port, config.MySqlConf.DbName)

	db0, err := gorm.Open("mysql", server)
	if err != nil {
		fmt.Println("db err: ", err)
	}
	db0.DB().SetMaxIdleConns(10)
	db0.LogMode(true)
	db = db0
	return db
}

// This function will create a temporarily database for running testing cases
func TestDBInit() *gorm.DB {
	server := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", config.MySqlConf.UserName, config.MySqlConf.PassWd, config.MySqlConf.Host, config.MySqlConf.Port, config.MySqlConf.DbName)
	test_db, err := gorm.Open("mysql", server)
	if err != nil {
		fmt.Println("db err: ", err)
	}
	test_db.DB().SetMaxIdleConns(3)
	test_db.LogMode(true)
	db = test_db
	return db
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return db
}
