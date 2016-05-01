package main

import (
	"consolia-api/web"
	"consolia-api/utils"
	_"fmt"
	"os"
)

func main() {
	conf := utils.Config{
		Env: os.Getenv("consolia_env"),
		DBHost: os.Getenv("consolia_db_host"),
		DBPort: os.Getenv("consolia_db_port"),
		Port: os.Getenv("consolia_port"),
	}

	server := web.NewServer(conf)
	server.Run(":" + conf.Port)
}