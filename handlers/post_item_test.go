package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Items.Store.Api/data"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
)

func TestPostItem(t *testing.T) {
	r := gin.Default()

	r.POST("/items", Create)

	// c.JSON(http.StatusOK, data.Item{})
	body, _ := json.Marshal(data.Item{
		UUID:        uuid.UUID{},
		Name:        "T",
		Price:       0,
		Description: "",
		SKU:         "",
		CreatedAt:   "",
		UpdatedAt:   "",
		DeletedAt:   "",
	})

	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, http.StatusCreated, w.Code)
}
