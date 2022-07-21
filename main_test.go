package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Items.Store.Api/handlers"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := SetUpRouter()
	router.GET("/items", handlers.GetAllItems)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
