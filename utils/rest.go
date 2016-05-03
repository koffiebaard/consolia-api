package utils

import (
	"gopkg.in/unrolled/render.v1"
	"net/http"
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
		"error": "Not found",
	})
}

func (r *Rest) InternalServerError(w http.ResponseWriter, error interface{}) {
	r.renderer.JSON(w, http.StatusInternalServerError, error)
}

func (r *Rest) OK(w http.ResponseWriter, data  interface{}) {
	r.renderer.JSON(w, http.StatusOK, data)
}
