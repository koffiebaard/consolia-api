package web

import (
	"github.com/codegangsta/negroni"
	"consolia-api/controllers"
	"consolia-api/utils"
	"gopkg.in/unrolled/render.v1"
	"github.com/gorilla/mux"
	_"fmt"
)

func NewServer (conf utils.Config) *negroni.Negroni {

	router := mux.NewRouter()
	renderer := render.New()

	// Comic
	comicController := controllers.NewComicController(conf, renderer)
	comicController.Register(router)

	// /etc
	etcController := controllers.NewEtcController(conf, renderer)
	etcController.Register(router)

	n := negroni.Classic()
	n.UseHandler(router)

	return n
}