package handlers

import (
	"net/http"

	"github.com/Items.Store.Api/services"
	"github.com/gin-gonic/gin"
)

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

	resp, _ := item_service.GetById(ctx, id)

	if resp != nil {
		ctx.JSON(http.StatusOK, resp)
	}

	ctx.JSON(http.StatusNotFound, nil)
}
