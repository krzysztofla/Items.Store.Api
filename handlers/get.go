package handlers

import (
	"net/http"

	"github.com/Items.Store.Api/data"
	"github.com/Items.Store.Api/db"
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
	repo, _ := db.NewRepository()

	repo.GetAllItems(ctx)

	ctx.JSON(http.StatusOK, data.Itmes{})
}
