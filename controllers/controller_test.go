package controllers

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "consolia-api/utils"
    "gopkg.in/unrolled/render.v1"
    "github.com/gorilla/mux"
    _"fmt"
)

func TestControllers(t *testing.T) {

    conf, err := utils.GetConf()
    assert.NoError(t, err)

    router := mux.NewRouter()
    renderer := render.New()
    REST := utils.GetSomeRest(renderer)

    assert.NotNil(t, router)
    assert.NotNil(t, renderer)
    assert.NotNil(t, REST)

    // Comic
    comicController := NewComicController(conf, renderer, conf.DB, REST)
    assert.NotNil(t, comicController)
    comicController.Register(router)

    // Comic history
    comicHistoryController := NewComicHistoryController(conf, renderer, conf.DB, REST)
    assert.NotNil(t, comicHistoryController)
    comicHistoryController.Register(router)
}
