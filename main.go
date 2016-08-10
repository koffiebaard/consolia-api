package main

import (
    "consolia-api/web"
	"consolia-api/utils"
	"fmt"
	_"os"
)

func main() {

    conf, err := utils.GetConf()

    if err != nil {
        fmt.Println(err)
    }

	server := web.NewServer(conf, conf.DB)
	server.Run(":" + conf.Port)
}
