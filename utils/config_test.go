package utils

import (
    _"fmt"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestGetDBConf(t *testing.T) {

    dbConf := GetDBConf()

    assert.IsType(t, dbConf, DBConfig{}, "DB configuration should be of type DBConfig")

    assert.NotNil(t, dbConf.DBHost, "DB host should be present in the DB configuration")
    assert.NotNil(t, dbConf.DBPort, "DB port should be present in the DB configuration")
    assert.NotNil(t, dbConf.DBUsername, "DB username should be present in the DB configuration")
    assert.NotNil(t, dbConf.DBPassword, "DB password should be present in the DB configuration")
    assert.NotNil(t, dbConf.DBName, "DB name should be present in the DB configuration")
}

func TestGetConf(t *testing.T) {

    conf, err := GetConf()

    assert.NoError(t, err)

    assert.IsType(t, conf, Config{}, "Exposed configuration should be of type Config")

    assert.NotNil(t, conf.Env, "Env should be present in the configuration")
    assert.NotNil(t, conf.Port, "Port should be present in the configuration")
}
