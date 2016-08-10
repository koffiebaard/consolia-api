package utils

import (
	"gopkg.in/unrolled/render.v1"
	"net/http"
    _ "fmt"
)

type Rest struct {
	renderer *render.Render
}

func GetSomeRest(renderer *render.Render) *Rest {

	return &Rest{
		renderer: renderer,
	}
}

func (r *Rest) NotFound(w http.ResponseWriter) {
	r.renderer.JSON(w, http.StatusNotFound, map[string]string{
		"message": "Aww. Not found.",
	})
}

func (r *Rest) InternalServerError(w http.ResponseWriter, error string) {
    r.renderer.JSON(w, http.StatusInternalServerError, map[string]string{
        "message": error,
    })
}

func (r *Rest) OK(w http.ResponseWriter, data interface{}) {
    r.renderer.JSON(w, http.StatusOK, data)
}


func (r *Rest) Unauthorized(w http.ResponseWriter, message string) {
    r.renderer.JSON(w, http.StatusUnauthorized, map[string]string{
        "message": message,
    })
}
