package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Items.Store.Api/data"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestGetAll(t *testing.T) {
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	r.GET("/items", GetAllItems)
	c.JSON(http.StatusOK, data.Item{})

	c.Request, _ = http.NewRequest("GET", "/items", nil)

	r.ServeHTTP(w, c.Request)

	var item data.Item

	c.ShouldBind(&item)

	assert.Equal(t, item.Name, "")
	assert.Equal(t, http.StatusOK, w.Code)
}
