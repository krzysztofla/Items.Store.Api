package handlers

import (
	"net/http"

	"github.com/Items.Store.Api/services"
	"github.com/gin-gonic/gin"
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

	resp := item_service.GetAll(ctx)

	ctx.JSON(http.StatusOK, resp)
}
