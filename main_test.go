package main

import (
    _"fmt"
    "testing"
    "consolia-api/utils"
    "consolia-api/web"
    "github.com/codegangsta/negroni"
    "github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {

    conf, err := utils.GetConf()

    assert.NoError(t, err)

    server := web.NewServer(conf, conf.DB)

    assert.IsType(t, server, negroni.Classic());
}
