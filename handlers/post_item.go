package handlers

import (
	"log"
	"net/http"

	"github.com/Items.Store.Api/data"
	"github.com/Items.Store.Api/services"
	"github.com/gin-gonic/gin"
)

// CreateItem godoc
// @Summary Create new item in database
// @Schemes
// @Description List items
// @Tags listAll
// @Accept json
// @Produce json
// @Success 200 {array} data.Item
// @Router /items [get]
func Create(ctx *gin.Context) {
	var item data.Item

	item_service, _ := services.NewItemsService()

	err := ctx.ShouldBind(&item)

	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	item_service.CreateItem(ctx, item)

	ctx.JSON(http.StatusOK, nil)
}
