package controllers

import (
    "consolia-api/utils"
    "consolia-api/models"
    "gopkg.in/unrolled/render.v1"
    "github.com/gorilla/mux"
    "github.com/jinzhu/gorm"
    "net/http"
    _"fmt"
)

type ComicHistoryController struct {
    conf utils.Config
    renderer *render.Render
    db *gorm.DB
    REST *utils.Rest
}

func NewComicHistoryController(conf utils.Config, renderer *render.Render, db *gorm.DB, REST *utils.Rest) *ComicHistoryController {

    return &ComicHistoryController{
        conf: conf,
        renderer: renderer,
        db: db,
        REST: REST,
    }
}

func (cc *ComicHistoryController) Register(router *mux.Router) {

    router.HandleFunc("/comic_history/", cc.listComicHistory).Methods("GET")
    router.HandleFunc("/comic_history/get_trending", cc.getTrending).Methods("GET")
}

func (cc *ComicHistoryController) listComicHistory (w http.ResponseWriter, r *http.Request) {

    // comic := models.Comic{}
    // comicHistory := comic.GetComicActivity(cc.db)
    cc.REST.OK(w, "stub")
}

func (cc *ComicHistoryController) getTrending (w http.ResponseWriter, r *http.Request) {

    comic := models.Comic{}
    trending := comic.GetTrending(cc.db)
    cc.REST.OK(w, trending)
}
