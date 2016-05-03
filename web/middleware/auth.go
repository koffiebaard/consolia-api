package middleware

import (
	"github.com/RangelReale/osin"
	"github.com/wisc/osin-mysql"
	"net/http"
	"github.com/jinzhu/gorm"
	_"fmt"
)


type AuthMiddleware struct {
	DB *gorm.DB
}

func (am *AuthMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	if r.RequestURI[0:len("/token")] != "/token" {
		auth := osin.CheckBearerAuth(r)

		if auth != nil {
			token := auth.Code
			store := mysql.New(am.DB.DB())
			_, err := store.LoadAccess(token)

			if err != nil {
				http.Error(rw, "Supplied access token is invalid", 401)
				return
			}
		} else {
			http.Error(rw, "No access token supplied", 401)
			return
		}
	}

	next(rw, r)
}
