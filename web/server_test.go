package web

import (
    _"fmt"
    "testing"
    "consolia-api/utils"
    "github.com/codegangsta/negroni"
    "github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {

    conf, err := utils.GetConf()

    assert.NoError(t, err)

    server := NewServer(conf, conf.DB)

    assert.IsType(t, server, negroni.New());

    handlers := server.Handlers()
    assert.Equal(t, 2, len(handlers), "Number of handlers should be 2 (1 middleware, 1 router)")
}
