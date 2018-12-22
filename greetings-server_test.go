package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/greeting", nil)
	router.ServeHTTP(w, req)

	json := `{"message":"hello, world"}`
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, json, w.Body.String())
}
