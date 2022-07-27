package handlers

import (
	"net/http"

	"github.com/Items.Store.Api/services"
	"github.com/gin-gonic/gin"
)

// DeleteItem godoc
// @Summary Delete item by unique id
// @Schemes
// @Description Delete item by id
// @Tags byId
// @Accept json
// @Produce json
// @Param id path string true "Item guid"
// @Success 200
// @Router /items/{id} [delete]
func Delete(ctx *gin.Context) {
	item_service, _ := services.NewItemsService()
	id := ctx.Param("id")

	item_service.DeleteItem(ctx, id)
	ctx.JSON(http.StatusNoContent, nil)
}
