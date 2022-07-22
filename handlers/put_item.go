package handlers

import (
	"log"
	"net/http"

	"github.com/Items.Store.Api/data"
	"github.com/Items.Store.Api/services"
	"github.com/gin-gonic/gin"
)

// UpdateItem godoc
// @Summary Updating item properties in database
// @Schemes
// @Description Update
// @Tags updateItem
// @Accept json
// @Produce json
// @Success 201 {string} uuid
// @Router /items [put]
func Update(ctx *gin.Context) {
	var item data.Item

	item_service, _ := services.NewItemsService()

	err := ctx.ShouldBind(&item)

	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	item_service.UpdateItem(ctx, item)

	ctx.JSON(http.StatusCreated, item.UUID)
}
