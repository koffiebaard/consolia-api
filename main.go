package main

import (
	"consolia-api/web"
	"consolia-api/utils"
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	conf := utils.Config{
		Env: 		os.Getenv("consolia_env"),
		DBHost: 	os.Getenv("consolia_db_host"),
		DBPort: 	os.Getenv("consolia_db_port"),
		DBUsername: os.Getenv("consolia_db_username"),
		DBPassword: os.Getenv("consolia_db_password"),
		DBName: 	os.Getenv("consolia_db_name"),
		Port: 		os.Getenv("consolia_port"),
	}

	db, err := gorm.Open("mysql", conf.DBUsername + ":" + conf.DBPassword + "@tcp(" + conf.DBHost + ":" + conf.DBPort + ")/" + conf.DBName + "?charset=utf8&parseTime=True")

    if err != nil {
		fmt.Printf("Got error when connecting to the database: '%v'", err)
	}

	if conf.Env == "dev" {
		db.LogMode(true)
	}

	server := web.NewServer(conf, db)
	server.Run(":" + conf.Port)
}
