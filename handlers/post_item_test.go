package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Items.Store.Api/data"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

func TestPostItem(t *testing.T) {
	r := gin.Default()

	r.POST("/items", CreateMock)

	// c.JSON(http.StatusOK, data.Item{})
	body, _ := json.Marshal(data.Item{
		Name:        "Test",
		Price:       1233,
		Description: "",
		SKU:         "",
	})

	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func CreateMock(ctx *gin.Context) {
	var item data.Item
	err := ctx.ShouldBind(&item)

	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	validate := validator.New()

	item.UUID = uuid.New()

	err = validate.Struct(&item)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		log.Println(validationErrors)
	}

	ctx.JSON(http.StatusCreated, nil)
}
