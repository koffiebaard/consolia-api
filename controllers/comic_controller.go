package controllers

import (
	"consolia-api/utils"
	"consolia-api/models"
	"gopkg.in/unrolled/render.v1"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"encoding/json"
	_"fmt"
)

type ComicController struct {
	conf utils.Config
	renderer *render.Render
	db *gorm.DB
	REST *utils.Rest
}

func NewComicController(conf utils.Config, renderer *render.Render, db *gorm.DB, REST *utils.Rest) *ComicController {

	return &ComicController{
		conf: conf,
		renderer: renderer,
		db: db,
		REST: REST,
	}
}

func (cc *ComicController) Register(router *mux.Router) {

	router.HandleFunc("/comic", cc.ListComics).Methods("GET")
	router.HandleFunc("/comic", cc.PostComic).Methods("POST")
	router.HandleFunc("/comic/{id:[0-9]+}", cc.GetComic).Methods("GET")
	router.HandleFunc("/comic/{id:[0-9]+}", cc.PatchComic).Methods("PATCH")
	router.HandleFunc("/comic/get_popular", cc.GetPopularSocialMedia).Methods("GET")
	router.HandleFunc("/awesome_stats", cc.GetAwesomeStats).Methods("GET")
	router.HandleFunc("/comic/get_upcoming_comic", cc.GetUpcomingComic).Methods("GET")
    router.HandleFunc("/comic/notifications", cc.GetNotifications).Methods("GET")
	router.HandleFunc("/comic/get_publish_activity", cc.GetPublishActivity).Methods("GET")
}

func (cc *ComicController) GetComic (w http.ResponseWriter, r *http.Request) {

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

func (cc *ComicController) ListComics (w http.ResponseWriter, r *http.Request) {
	comics := []models.Comic{}
	cc.db.Find(&comics)

	cc.REST.OK(w, comics)
}


func (cc *ComicController) PatchComic (w http.ResponseWriter, r *http.Request) {
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

func (cc *ComicController) PostComic (w http.ResponseWriter, r *http.Request) {

	// Decode req body
	decoder := json.NewDecoder(r.Body)
	var newComic models.Comic
	err := decoder.Decode(&newComic)
	if err != nil {
		cc.REST.InternalServerError(w, "The request body is not formatted correctly. " + err.Error())
        return
	}

	// Update the data
	if err := cc.db.Model(&newComic).Save(&newComic).Error; err != nil {
		cc.REST.InternalServerError(w, "Validation error. " + err.Error())
		return
	}

	// Show to user
	cc.REST.OK(w, newComic)
}


func (cc *ComicController) GetPopularSocialMedia (w http.ResponseWriter, r *http.Request) {

	comic := models.Comic{}
	comics := comic.GetPopularSocialMedia(cc.db)

	cc.REST.OK(w, comics)
}

func (cc *ComicController) GetAwesomeStats (w http.ResponseWriter, r *http.Request) {

    comic := models.Comic{}
    awesomeStats := comic.GetAwesomeStats(cc.db)
    cc.REST.OK(w, awesomeStats)
}

func (cc *ComicController) GetUpcomingComic (w http.ResponseWriter, r *http.Request) {

    comic := models.Comic{}
    upcomingComic := comic.GetUpcomingComic(cc.db)
	cc.REST.OK(w, upcomingComic)
}

func (cc *ComicController) GetNotifications (w http.ResponseWriter, r *http.Request) {

    comic := models.Comic{}
    notifications := comic.GetNotifications(cc.db)
    cc.REST.OK(w, notifications)
}

func (cc *ComicController) GetPublishActivity (w http.ResponseWriter, r *http.Request) {

    comic := models.Comic{}
    comicActivities := comic.GetPublishActivity(cc.db)
    cc.REST.OK(w, comicActivities)
}
