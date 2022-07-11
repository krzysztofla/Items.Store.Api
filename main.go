package main

import (
	"github.com/Items.Store.Api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/items", handlers.GetAllItems)
	r.Run()
}
