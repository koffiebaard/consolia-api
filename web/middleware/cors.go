package middleware

import (
    "net/http"
    "github.com/jinzhu/gorm"
    _"fmt"
)


type CorsMiddleware struct {
    DB *gorm.DB
}

func (am *CorsMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

    rw.Header().Set("Access-Control-Allow-Origin", "*")
    rw.Header().Set("Access-Control-Max-Age", "1728000")
    rw.Header().Set("Access-Control-Allow-Headers", "Authorization")

    next(rw, r)
}
