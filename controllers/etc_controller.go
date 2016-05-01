package controllers

import (
	"consolia-api/utils"
	"gopkg.in/unrolled/render.v1"
	"net/http"
	"github.com/gorilla/mux"
)

type EtcController struct {
	conf utils.Config
	renderer *render.Render
}

func NewEtcController(conf utils.Config, renderer *render.Render) *EtcController {

	return &EtcController{
		conf: conf,
		renderer: renderer,
	}
}

func (cc *EtcController) Register(router *mux.Router) {

	router.HandleFunc("/", cc.home).Methods("GET")
}

func (cc *EtcController) home (w http.ResponseWriter, r *http.Request) {

	cc.renderer.JSON(w, http.StatusOK, "Good day, dear sir.")
}