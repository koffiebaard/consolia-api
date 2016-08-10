package web

import (
	"github.com/codegangsta/negroni"
	"consolia-api/controllers"
	"consolia-api/utils"
	"consolia-api/web/middleware"
	"gopkg.in/unrolled/render.v1"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_"fmt"
    "net/http"
)

func NewServer (conf utils.Config, db *gorm.DB) *negroni.Negroni {

	router := mux.NewRouter()
	renderer := render.New()
	REST := utils.GetSomeRest(renderer)

	// Comic
    comicController := controllers.NewComicController(conf, renderer, db, REST)
    comicController.Register(router)

    // Comic history
	comicHistoryController := controllers.NewComicHistoryController(conf, renderer, db, REST)
	comicHistoryController.Register(router)

	// Auth
	authController := controllers.NewAuthController(conf, renderer, db, REST)
	authController.Register(router)

	// /etc
	etcController := controllers.NewEtcController(conf, renderer, db, REST)
	etcController.Register(router)

    router.NotFoundHandler = http.HandlerFunc(notFound)

    n := negroni.New()
    n.Use(&middleware.AuthMiddleware{DB: db})
    n.UseHandler(router)

	return n
}

func notFound(rw http.ResponseWriter, r *http.Request) {

    renderer := render.New()
    REST := utils.GetSomeRest(renderer)
    REST.NotFound(rw)
}
