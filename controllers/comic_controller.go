package controllers

import (
	"consolia-api/utils"
	"gopkg.in/unrolled/render.v1"
	"github.com/gorilla/mux"
	"net/http"
)

type ComicController struct {
	conf utils.Config
	renderer *render.Render
}

func NewComicController(conf utils.Config, renderer *render.Render) *ComicController {

	return &ComicController{
		conf: conf,
		renderer: renderer,
	}
}

func (cc *ComicController) Register(router *mux.Router) {

	router.HandleFunc("/comic/{id:[0-9]+}", cc.getComic).Methods("GET")
}

func (cc *ComicController) getComic (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	cc.renderer.JSON(w, http.StatusOK, map[string]string{
		"error": "lel j/k",
		"id": id,
	})
}