// +build main2
package main

import (
    "github.com/mattes/migrate/migrate"
    _ "github.com/mattes/migrate/driver/mysql"
    "consolia-api/utils"
    "os"
    "fmt"
)

func main() {

    conf := utils.GetDBConf()

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
