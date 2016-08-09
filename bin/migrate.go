package main

import "github.com/mattes/migrate/migrate"
import _ "github.com/mattes/migrate/driver/mysql"
import "consolia-api/utils"
import "os"
import "fmt"

func main() {

    conf := utils.Config{
            Env:        os.Getenv("consolia_env"),
            DBHost:     os.Getenv("consolia_db_host"),
            DBPort:     os.Getenv("consolia_db_port"),
            DBUsername: os.Getenv("consolia_db_username"),
            DBPassword: os.Getenv("consolia_db_password"),
            DBName:     os.Getenv("consolia_db_name"),
            Port:       os.Getenv("consolia_port"),
        }

    if len(os.Args) == 2 && os.Args[1] == "up" {
        allErrors, ok := migrate.UpSync("mysql://" + conf.DBUsername + ":" + conf.DBPassword + "@tcp(" + conf.DBHost + ":" + conf.DBPort + ")/" + conf.DBName + "?charset=utf8&parseTime=True", "./migrations")

        if !ok {
            fmt.Println(allErrors)
        }
    } else if len(os.Args) == 2 && os.Args[1] == "down" {
        allErrors, ok := migrate.DownSync("mysql://" + conf.DBUsername + ":" + conf.DBPassword + "@tcp(" + conf.DBHost + ":" + conf.DBPort + ")/" + conf.DBName + "?charset=utf8&parseTime=True", "./migrations")

        if !ok {
            fmt.Println(allErrors)
        }
    } else {
        fmt.Println("No valid argument entered. 'up' or 'down'.");
    }
}
