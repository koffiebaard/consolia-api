package main

import (
    _"fmt"
    "testing"
    "consolia-api/utils"
    "github.com/mattes/migrate/migrate"
    "github.com/stretchr/testify/assert"
    _ "github.com/mattes/migrate/driver/mysql"
)

func TestMigrate(t *testing.T) {

    conf := utils.GetDBConf()

    assert.IsType(t, conf, utils.DBConfig{}, "DB configuration should be of type DBConfig")

    _, err := migrate.Version("mysql://" + conf.DBUsername + ":" + conf.DBPassword + "@tcp(" + conf.DBHost + ":" + conf.DBPort + ")/" + conf.DBName + "?charset=utf8&parseTime=True", "./migrations")

    assert.NoError(t, err)
}
