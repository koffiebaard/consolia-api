package middleware

import (
	"github.com/RangelReale/osin"
    "github.com/wisc/osin-mysql"
    "consolia-api/utils"
	"net/http"
	"github.com/jinzhu/gorm"
	_"fmt"
    "gopkg.in/unrolled/render.v1"
)


type AuthMiddleware struct {
	DB *gorm.DB
}

func (am *AuthMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	if r.RequestURI != "/token" {
		auth := osin.CheckBearerAuth(r)

		if auth != nil {
			token := auth.Code
			store := mysql.New(am.DB.DB())
			_, err := store.LoadAccess(token)

            renderer := render.New()
            REST := utils.GetSomeRest(renderer)

			if err != nil {
                REST.Unauthorized(rw, "Supplied access token is invalid")
                return
            }
        } else {
            REST.Unauthorized(rw, "No access token supplied")
			return
		}
	}

	next(rw, r)
}
