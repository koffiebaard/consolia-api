package controllers

import (
	"consolia-api/utils"
	"consolia-api/models"
	"gopkg.in/unrolled/render.v1"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"github.com/RangelReale/osin"
	"github.com/wisc/osin-mysql"
)

type AuthController struct {
	conf utils.Config
	renderer *render.Render
	db gorm.DB
	REST *utils.Rest
}

func NewAuthController(conf utils.Config, renderer *render.Render, db *gorm.DB, REST *utils.Rest) *AuthController {

	return &AuthController{
		conf: conf,
		renderer: renderer,
		db: *db,
		REST: REST,
	}
}

func (cc *AuthController) Register(router *mux.Router) {

	router.HandleFunc("/token/", cc.token).Methods("POST")
}

func (cc *AuthController) token (w http.ResponseWriter, r *http.Request) {

	store := mysql.New(cc.db.DB())
	config := osin.NewServerConfig()
	config.AllowedAccessTypes = osin.AllowedAccessType{osin.PASSWORD, osin.CLIENT_CREDENTIALS}
	config.ErrorStatusCode = 401
	server := osin.NewServer(config, store)

	resp := server.NewResponse()
	defer resp.Close()

	if ar := server.HandleAccessRequest(resp, r); ar != nil {

		user := models.User{}

		if user.Authenticate(cc.db, ar.Username, ar.Password) {
			ar.Authorized = true
		}

		server.FinishAccessRequest(resp, r, ar)
	}

	osin.OutputJSON(resp, w, r)
}