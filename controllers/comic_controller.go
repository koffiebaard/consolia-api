package controllers

import (
	"consolia-api/utils"
	"consolia-api/models"
	"gopkg.in/unrolled/render.v1"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"encoding/json"
)

type ComicController struct {
	conf utils.Config
	renderer *render.Render
	db gorm.DB
	REST *utils.Rest
}

func NewComicController(conf utils.Config, renderer *render.Render, db *gorm.DB, REST *utils.Rest) *ComicController {

	return &ComicController{
		conf: conf,
		renderer: renderer,
		db: *db,
		REST: REST,
	}
}

func (cc *ComicController) Register(router *mux.Router) {

	router.HandleFunc("/comic/", cc.listComics).Methods("GET")
	router.HandleFunc("/comic/", cc.postComic).Methods("POST")
	router.HandleFunc("/comic/{id:[0-9]+}", cc.getComic).Methods("GET")
	router.HandleFunc("/comic/{id:[0-9]+}", cc.patchComic).Methods("PATCH")
}

func (cc *ComicController) getComic (w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	comic := models.Comic{}

	// Catch 404
	if cc.db.First(&comic, id).Error != nil {
		cc.REST.NotFound(w)
		return
	}

	cc.REST.OK(w, comic)
}

func (cc *ComicController) listComics (w http.ResponseWriter, r *http.Request) {
	comics := []models.Comic{}
	cc.db.Find(&comics)

	cc.REST.OK(w, comics)
}


func (cc *ComicController) patchComic (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	// Make sure ID exists
	comic := models.Comic{}
	if cc.db.First(&comic, id).Error != nil {
		cc.REST.NotFound(w)
	}

	// Decode req body
	decoder := json.NewDecoder(r.Body)
	var updatedComic models.Comic
	err := decoder.Decode(&updatedComic)
	if err != nil {
		cc.REST.InternalServerError(w, err.Error())
	}

	// Update the data
	if err := cc.db.Model(&comic).Updates(&updatedComic).Error; err != nil {
		cc.REST.InternalServerError(w, err.Error())
		return
	}

	// Show to user
	cc.REST.OK(w, comic)
}

func (cc *ComicController) postComic (w http.ResponseWriter, r *http.Request) {

	// Decode req body
	decoder := json.NewDecoder(r.Body)
	var newComic models.Comic
	err := decoder.Decode(&newComic)
	if err != nil {
		cc.REST.InternalServerError(w, err.Error())
	}

	// Update the data
	if err := cc.db.Model(&newComic).Save(&newComic).Error; err != nil {
		cc.REST.InternalServerError(w, err.Error())
		return
	}

	// Show to user
	cc.REST.OK(w, newComic)
}