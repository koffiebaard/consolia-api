package middleware

import (
    _"fmt"
    "testing"
    "consolia-api/utils"
    "github.com/stretchr/testify/assert"
    "gopkg.in/unrolled/render.v1"
)

func TestREST(t *testing.T) {

    renderer := render.New()
    REST := utils.GetSomeRest(renderer)

    assert.NotNil(t, REST)
}
