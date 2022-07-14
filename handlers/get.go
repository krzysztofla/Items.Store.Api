package handlers

import (
	"net/http"

	"github.com/Items.Store.Api/data"
	"github.com/Items.Store.Api/services"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// GetAllItems
// @Summary Gets all items from database
// @Schemes
// @Description
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {array} data.Item
// @Router /items [get]
func GetAllItems(ctx *gin.Context) {
	item_service, _ := services.NewItemsService()

	item_service.GetAll(ctx)

	ctx.JSON(http.StatusOK, data.Itmes{})
}
