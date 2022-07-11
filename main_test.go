package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Items.Store.Api/handlers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := gin.Default()
	router.GET("/items", handlers.GetAllItems)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
