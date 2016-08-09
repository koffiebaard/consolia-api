package controllers

import (
	"consolia-api/utils"
	"gopkg.in/unrolled/render.v1"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

type EtcController struct {
	conf utils.Config
	renderer *render.Render
	db gorm.DB
	REST *utils.Rest
}

func NewEtcController(conf utils.Config, renderer *render.Render, db *gorm.DB, REST *utils.Rest) *EtcController {

	return &EtcController{
		conf: conf,
		renderer: renderer,
		db: *db,
		REST: REST,
	}
}

func (cc *EtcController) Register(router *mux.Router) {

    router.HandleFunc("/", cc.home).Methods("GET")
	router.HandleFunc("/health", cc.health).Methods("GET")
}

func (cc *EtcController) home (w http.ResponseWriter, r *http.Request) {

	cc.REST.OK(w, "Good day, dear sir.")
}

func (cc *EtcController) health (w http.ResponseWriter, r *http.Request) {

    cc.REST.OK(w, "OK")
}
