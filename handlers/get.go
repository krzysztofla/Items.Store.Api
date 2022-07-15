package handlers

import (
	"net/http"

	"github.com/Items.Store.Api/data"
	"github.com/Items.Store.Api/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetAllItems godoc
// @Summary Gets all items from database
// @Schemes
// @Description List items
// @Tags listAll
// @Accept json
// @Produce json
// @Success 200 {array} data.Item
// @Router /items [get]
func GetAllItems(ctx *gin.Context) {
	item_service, _ := services.NewItemsService()

	item_service.GetAll(ctx)

	ctx.JSON(http.StatusOK, data.Itmes{})
}

// GetItemById godoc
// @Summary Get item by unique id
// @Schemes
// @Description Get item by id
// @Tags byId
// @Accept json
// @Produce json
// @Param id path string true "Item guid"
// @Success 200 {object} data.Item
// @Router /items/{id} [get]
func GetItemById(ctx *gin.Context) {
	item_service, _ := services.NewItemsService()
	id := ctx.Param("id")

	uid, _ := uuid.Parse(id)

	item_service.GetById(ctx, uid)
}
