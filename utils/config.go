package utils

import (
    "os"
    _"fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
    DBHost string
    DBPort string
    DBUsername string
    DBPassword string
    DBName string
}

type Config struct {
    Env string
    Port string
    DB *gorm.DB
}

func GetDBConf() (DBConfig) {
    dbConf := DBConfig{
        DBHost:     os.Getenv("consolia_db_host"),
        DBPort:     os.Getenv("consolia_db_port"),
        DBUsername: os.Getenv("consolia_db_username"),
        DBPassword: os.Getenv("consolia_db_password"),
        DBName:     os.Getenv("consolia_db_name"),
    }

    return dbConf
}

func GetConf() (Config, error) {

    dbConf := GetDBConf()

    exposeConf := Config{
        Env:        os.Getenv("consolia_env"),
        Port:       os.Getenv("consolia_port"),
    }

    db, err := gorm.Open("mysql", dbConf.DBUsername + ":" + dbConf.DBPassword + "@tcp(" + dbConf.DBHost + ":" + dbConf.DBPort + ")/" + dbConf.DBName + "?charset=utf8&parseTime=True")

    if err != nil {
        return exposeConf, err
    }

    if exposeConf.Env == "dev" {
        db.LogMode(true)
    }

    exposeConf.DB = db

    return exposeConf, nil
}
